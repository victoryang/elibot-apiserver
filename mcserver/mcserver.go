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

var Mcs *MCserver = nil

func GetMcServer() *MCserver {
	return Mcs
}

func (mc *MCserver) Exec(cmd string, from string, rch chan Response) {
	Log.Debug("MCserver exec ", cmd, " from ", from)
	mc.WorkChan <- Request{
		Command: cmd, 
		From: from, 
		RespCh: rch,
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