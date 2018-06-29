package mcserver

import (
	"errors"
	"net"
	"context"

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

func getConnFromPool() (interface{}, Response) {
	Log.Debug("MCServer get a connection from pool...")
	if Mcs == nil {
		return nil, Response{Result: "", Err: errors.New("MCServer Error")}
	}

	conn, err := Mcs.ConnPool.Get()
	if err!=nil {
		return nil, Response{Result: "", Err: err}
	}
	return conn, Response{}
}

func execute(ctx context.Context, ch chan Response, cmd string) {
	Log.Debug("MCServer executing new command...")
	conn, resp := getConnFromPool()
	if conn == nil {
		SafeSendResponseToChannel(ch, resp)
		return
	}
	defer func() {
		Mcs.ConnPool.Put(conn)
	}()

	select {
	case <-ctx.Done():
		Log.Debug("job cancelled")
		SafeSendResponseToChannel(ch, Response{Result: "", Err: errors.New("MCserver job cancelled")})
		return
	default:
		res, err := HandleCommand(conn.(net.Conn), cmd)
		SafeSendResponseToChannel(ch, Response{Result: res, Err: err})
	}
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case req := <- Mcs.WorkChan:
			go execute(ctx, req.RespCh, req.Command)
		}
	}
}

func (mc *MCserver) Close() {
	if mc!=nil {
		close(mc.WorkChan)
		mc.Cancel()
		mc.ConnPool.Release()
		mc=nil
		Log.Print("MCServer closed")
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
		WorkChan:   make(chan Request),
		Ctx: 		ctx,
		Cancel:		cancel,
	}
	Log.Print("MCServer started")
	go worker(Mcs.Ctx)
	return Mcs
}