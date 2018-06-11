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

func OnCommandRecived(ctx context.Context, command string) (res string, err error) {
	if mcserver == nil {
		return "", errors.New("MCServer Error")
	}

	var connection interface{}
	connection, err = mcserver.ConnPool.Get()
	if err!=nil {
		res = ""
		Log.Error("MCServer error: can not get a connection now, try it again later")
		return
	}
	defer mcserver.ConnPool.Put(connection)

	done := make(chan bool)
	go func(c chan bool, conn interface{}, cmd string) {
		if consumeCommandLineIf(conn.(net.Conn)) {
			res, err = handleCommand(conn.(net.Conn), cmd)
		} else {
			err = errors.New("MCServer error: bad connection")
			res = ""
		}
		c<-true
	}(done, connection, command)

	DONE:
	for {
		select {
		case <-ctx.Done():
			err = errors.New("MCServer: context done")
			break DONE
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