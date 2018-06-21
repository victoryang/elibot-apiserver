package mcserver

import (
	"errors"
	"net"
	"context"
	"time"

	Log "elibot-apiserver/log"
	"elibot-apiserver/mcserver/pool"
)

type MCserver struct {
	Address			string
	ConnPool 		pool.Pool
	WorkChan		chan Request
	Ctx 			context.Context
	Cancel 			context.CancelFunc
}

type Request struct {
	Command 		string
	From			string
	RespCh 			chan Response
}

type Response struct {
	Result 			string
	Err 			error
}

var Mcs *MCserver = nil
const (
	MaxJobExecDuration = 200
)

func handleCommand(conn net.Conn, command string) (string, error) {
	err := WriteMessage(conn, command)
	if err!=nil {
		return "", err
	}

	return ReadMessage(conn)
}

func getConnFromPool() (interface{}, Response) {
	if Mcs == nil {
		return nil, Response{Result: "", Err: errors.New("MCServer Error")}
	}

	conn, err := Mcs.ConnPool.Get()
	if err!=nil {
		return nil, Response{Result: "", Err: errors.New("MCServer error: can not get a connection now, try it again later")}
	}
	return conn, Response{}
}

func execute(ctx context.Context, ch chan Response, cmd string) {
	conn, resp := getConnFromPool()
	if conn == nil {
		SafeSendResponseToChannel(ch, resp)
		return
	}
	defer func() {
		if _, ok := <- ch; ok {
	        close(ch)
	    }
		Mcs.ConnPool.Put(conn)
	}()

	select {
	case <-ctx.Done():
		SafeSendResponseToChannel(ch, Response{Result: "", Err: errors.New("MCserver job cancelled")})
		return
	default:
		if consumeCommandLineIf(conn.(net.Conn)) {
			res, err := handleCommand(conn.(net.Conn), cmd)
			SafeSendResponseToChannel(ch, Response{Result: res, Err: err})
		} else {
			SafeSendResponseToChannel(ch, Response{Result: "", Err: errors.New("MCServer error: bad connection")})
		}
	}
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case req := <- Mcs.WorkChan:
			/* pass ctx to all child for graceful shutdown*/
			ctx_job, cancel := context.WithTimeout(ctx, MaxJobExecDuration * time.Millisecond)
			defer cancel()
			go execute(ctx_job, req.RespCh, req.Command)
		}
	}
}

func (mc *MCserver) Close() {
	if mc!=nil {
		mc.Cancel()
		close(mc.WorkChan)
		mc.ConnPool.Release()
		mc=nil
		Log.Debug("MCServer closed")
	}
} 

func (mc *MCserver) Exec(cmd string, from string, rch chan Response) {
	Log.Debug("MCserver exec ", cmd, " from ", from)
	mc.WorkChan <- Request{
		Command: cmd, 
		From: from, 
		RespCh: rch,
	}
}

func GetMcServer() *MCserver {
	return Mcs
}

func NewMCServer(address string, cap int) *MCserver {
	factory := func() (interface{}, error){return net.Dial("tcp", address)}
	close := func(v interface{}) (error){return v.(net.Conn).Close()}

	poolConfig := &pool.PoolConfig{
		InitialCap: cap,
		MaxCap:     5,
		Factory:    factory,
		Close:      close,
	}

	p, err := pool.NewChannelPool(poolConfig)
	if err!=nil {
		Log.Error("MCServer error: ", err)
		return Mcs
	}

	ctx, cancel := context.WithCancel(context.Background())
	Mcs = &MCserver{
		Address: 	address,
		ConnPool: 	p,
		WorkChan:   make(chan Request, 3),
		Ctx: 		ctx,
		Cancel:		cancel,
	}
	go worker(Mcs.Ctx)
	return Mcs
}