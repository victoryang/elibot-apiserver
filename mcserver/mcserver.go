package mcserver

import (
	"errors"
	"net"
	"context"
	"sync"

	Log "elibot-apiserver/log"
	"elibot-apiserver/mcserver/pool"
)

/*var test = "testGo 0 1"
var setzeroencode = "setZeroEncode 0 1"
var end = "\n"*/

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

var mcserver *MCserver = nil
var wg sync.WaitGroup

func handleCommand(conn net.Conn, command string) (string, error) {
	err := WriteMessage(conn, command)
	if err!=nil {
		return "", err
	}

	return ReadMessage(conn)
}

func getConnFromPool(ch chan Response) interface{} {
	if mcserver == nil {
		ch <- Response{Result: "", Err: errors.New("MCServer Error")}
		return nil
	}

	conn, err := mcserver.ConnPool.Get()
	if err!=nil {
		ch <- Response{Result: "", Err: errors.New("MCServer error: can not get a connection now, try it again later")}
		return nil
	}
	return conn
}

func execute(ctx context.Context, ch chan Response, conn interface{}, cmd string) {
	select {
	case <-ctx.Done():
		return
	default:
		if conn == nil {
			return
		}
		defer mcserver.ConnPool.Put(conn)
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

func worker(quit chan bool) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var req Request
	for {
		select {
		case req := <- mcserver.WorkChan:
			conn := getConnFromPool(req.Resp)
			go execute(ctx, req.Resp, conn, req.Command)
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
		return mcserver
	}

	mcserver = &MCserver{
		Address: 	address,
		ConnPool: 	p,
	}
	go worker(mcserver.QuitChan)
	return mcserver
}