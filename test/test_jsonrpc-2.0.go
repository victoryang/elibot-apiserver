package main

import (
	"net"
	"context"
	"fmt"
	"elibot-apiserver/jsonrpc2"
)

var JsonRpcClient *jsonrpc2.Conn
var address string = "192.168.1.166:8055"
var ctx, _ = context.WithCancel(context.Background())

type handler struct {
}

func (h *handler)Handle(ctx context.Context, c *jsonrpc2.Conn, req *jsonrpc2.Request) {
	fmt.Println("Receive request from server")
}

func testRequest() {
	var reply bool
	if err := JsonRpcClient.Call(ctx, "servo", []string{"on"}, &reply); err!=nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(reply)
}

func testRequest1() {
	var reply bool
	if err := JsonRpcClient.Call(ctx, "setArcParam",[]string{"arc.welder.prepareAspirationTime", "2", "0", "0"}, &reply); err!=nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(reply)
}

func main() {
	var err error
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Dial: ", err)
		return
	}

	stream := jsonrpc2.NewBufferedStream(conn, jsonrpc2.VarintObjectCodec{})
	JsonRpcClient = jsonrpc2.NewConn(ctx, stream, new(handler))

	testRequest()
	return
}