package mcserver

import (
	"errors"
	"net"
	"time"
	"context"

	Log "elibot-apiserver/log"
	"elibot-apiserver/mcserver/pool"
)

/*var test = "testGo 0 1"
var setzeroencode = "setZeroEncode 0 1"
var end = "\n"*/

type MCserver struct {
	Address			string
	ConnPool 		pool.Pool
}

var mcserver *MCserver = nil

func handleCommand(conn net.Conn, command string) (string, error) {
	err := WriteMessage(conn, command)
	if err!=nil {
		return "", err
	}

	return ReadMessage(conn)
}

func OnCommandRecived(ctx context.Context, cmd string) (res string, err error) {
	if mcserver == nil {
		return "", errors.New("MCServer Error")
	}

	var conn interface{}
	conn, err = mcserver.ConnPool.Get()
	if err!=nil {
		res = ""
		Log.Error("MCServer error: can not get a connection now, try it again later")
		return
	}
	defer mcserver.ConnPool.Put(conn)

	done := make(chan bool)
	go func(c chan bool) {
		if consumeCommandLineIf(conn.(net.Conn)) {
			res, err =  handleCommand(conn.(net.Conn), cmd)
		} else {
			err = errors.New("MCServer error: bad connection")
			res = ""
		}
		c<-true
	}(done)

	DONE:
	for {
		select {
		case <-ctx.Done():
			err = errors.New("MCServer: context done")
		case <-done:
			break DONE
		}
	}
	return
}

func (mc *MCserver) Close() {
	if mc!=nil {
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
		//链接最大空闲时间，超过该时间的链接 将会关闭，可避免空闲时链接EOF，自动失效的问题
		IdleTimeout: 15 * time.Second,
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
	return mcserver
}