package api

import (
	"net"
	"fmt"
	v2rpc "elibot-apiserver/api/v2"
)

const (
	port = 9500
)

type GrpcServer struct {
	grpc 		*grpc.Server
}

func (s *GrpcServer)Shutdown() {
	shutdownNow := func() {
		// close grpc.Server; cancels all active RPCs
		ss.grpc.Stop()
	}

	ch := make(chan struct{})
	go func() {
		defer close(ch)
		// close listeners to stop accepting new connections,
		// will block on any existing transports
		ss.grpc.GracefulStop()
	}()

	// wait until all pending RPCs are finished
	select {
	case <-ch:
		// took too long, manually close open transports
		// e.g. watch streams
		shutdownNow()

		// concurrent GracefulStop should be interrupted
		<-ch
	}
}

func NewGrpcServer() (*GrpcServer){
	lis, err := net.Listen("tcp", ":2500")
    if err != nil {
            fmt.Println("failed to listen: %v", err)
    }
    gs := v2rpc.Server()
    go gs.Serve(lis)
    return &GrpcServer{grpc: gs}
}