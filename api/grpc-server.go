package api

import (
	"net"
	"context"
	"time"

	Log "elibot-apiserver/log"
	"elibot-apiserver/config"

	v2rpc "elibot-apiserver/api/v2"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	grpc 		*grpc.Server
}

func (s *GrpcServer)Shutdown() {
	shutdownNow := func() {
		// close grpc.Server; cancels all active RPCs
		Log.Print("Grpc server: shutting down...")
		s.grpc.Stop()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	go func() {
		// close listeners to stop accepting new connections,
		// will block on any existing transports
		select {
		case <-ctx.Done():
			shutdownNow()
			return
		default:
			s.grpc.GracefulStop()
		}
	}()
}

func NewGrpcServer(c *config.GrpcEntryPoint) (*GrpcServer){
	lis, err := net.Listen("tcp", c.ListenAddress)
    if err != nil {
            Log.Error("failed to listen: %v", err)
            return nil
    }
    gs := v2rpc.Server(c)
    go gs.Serve(lis)

    Log.Print("Grpc server started...")

    return &GrpcServer{grpc: gs}
}