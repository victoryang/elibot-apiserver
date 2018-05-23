package main 

import (
	"net"
	"fmt"
)

func main() {
	conns := make([]net.Conn, 6)
	var conn net.Conn
	l, _ := net.Listen("tcp", ":9200")
	fmt.Println("start to listen...")
	for {
		conn, _ = l.Accept()
		conns = append(conns, conn)
		fmt.Println("new connection comming: have ", len(conns), "connections now")
	}
}