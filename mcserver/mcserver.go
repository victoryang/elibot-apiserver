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
	QuitChan		chan bool
}

type Request struct {
	Command 		string
	From			string
	Resp 			chan Response
}

type Response struct {
	Result 			string
	Err 			error
}

var Mcs *MCserver = nil

func handleCommand(conn net.Conn, command string) (string, error) {
	err := WriteMessage(conn, command)
	if err!=nil {
		return "", err
	}

	return ReadMessage(conn)
}

func getConnFromPool(ch chan Response) interface{} {
	if Mcs == nil {
		ch <- Response{Result: "", Err: errors.New("MCServer Error")}
		return nil
	}

	conn, err := Mcs.ConnPool.Get()
	if err!=nil {
		ch <- Response{Result: "", Err: errors.New("MCServer error: can not get a connection now, try it again later")}
		return nil
	}
	return conn
}

func execute(ctx context.Context, ch chan Response, cmd string) {
	for {
		select {
		case <-ctx.Done():
			ch<-Response{Result: "", Err: errors.New("MCserver job cancelled")}
			return
		default:
			conn := getConnFromPool(ch)
			if conn == nil {
				return
			}
			defer Mcs.ConnPool.Put(conn)
			var res string
			var err error
			if consumeCommandLineIf(conn.(net.Conn)) {
				res, err = handleCommand(conn.(net.Conn), cmd)
			} else {
				err = errors.New("MCServer error: bad connection")
				res = ""
			}
			ch<-Response{Result: res, Err: err}
		}
	}
}

func worker(quit chan bool) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		select {
		case req := <- Mcs.WorkChan:
			/* pass ctx to all child for graceful shutdown*/
			go execute(ctx, req.Resp, req.Command)
		case <-quit:
			return
		}
	}
}

func (mc *MCserver) Close() {
	if mc!=nil {
		mc.QuitChan<-true
		mc.ConnPool.Release()
		mc=nil
	}
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

	Mcs = &MCserver{
		Address: 	address,
		ConnPool: 	p,
		WorkChan:   make(chan Request),
		QuitChan: 	make(chan bool),
	}
	go worker(Mcs.QuitChan)
	return Mcs
}