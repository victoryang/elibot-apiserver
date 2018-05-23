package test

import (
    "testing"
    "fmt"
    "time"
    "net"

    "elibot-apiserver/mcserver/pool"
)

const (
	address = 127.0.0.1:9200
)

func Test_Pool(t *testing.T) {
	q := make(chan bool)
	conns := make([]*net.TCPConn)
	go func(quit chan bool){
		ln, _ := net.Listen("tcp", ":9200")
		go func(l net.Listener) {
			for {
				conn, err := l.Accept()
				conns = append(conns, conn)
				fmt.Println("new connection comming: have ", len(conns), "connections now")
			}
		}(ln)
		
		<-quit
	}(q)
	factory := func() {return net.Dial("tcp", address)}
	close := func(v interface{}) {return v.(net.Conn).Close()}

	poolConfig := &pool.PoolConfig{
		InitialCap: 5,
		MaxCap:     30,
		Factory:    factory,
		Close:      close,
		//链接最大空闲时间，超过该时间的链接 将会关闭，可避免空闲时链接EOF，自动失效的问题
		IdleTimeout: 15 * time.Second,
	}

	p, err := pool.NewChannelPool(poolConfig)
	if err != nil {
		fmt.Println("err=", err)
	}

	v, err := p.Get()
	p.Put(v)
	p.Release()
	current := p.Len()

	q<-true
}