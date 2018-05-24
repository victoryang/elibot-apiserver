package mcserver

import (
	"errors"
	"fmt"
	"net"

	Log "elibot-apiserver/log"
	"elibot-apiserver/mcserver/pool"
)

var test = "testGo 0 1"
var end = "\n"

type MCserver struct {
	Addr			string
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

func OnCommandRecived() (string, error) {
	cmd := test + end
	conn, err := mcserver.ConnPool.Get()
	if err!=nil {
		Log.Error("MCServer error: can not get a connection now, try it again later")
		return "", error
	}
	defer mcserver.ConnPool.Put(conn)

	var resp string
	if testCommandLine(conn) {
		return handleCommand(conn, cmd)
	} else {
		return "", errors.New("MCServer error: not in a normal situation")
	}
}

func (mc *MCserver) Close() {
	mc.ConnPool.Release()
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