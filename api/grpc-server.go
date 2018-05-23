package api

import (
	"net"
	"fmt"
	v2rpc "elibot-apiserver/api/v2"
)

const (
	port = 9500
)

func NewGrpcServer() {
	lis, err := net.Listen("tcp", ":2500")
    if err != nil {
            fmt.Println("failed to listen: %v", err)
    }
    gs := v2rpc.Server()
    go gs.Serve(lis)
}