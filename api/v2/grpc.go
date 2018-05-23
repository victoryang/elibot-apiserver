package v2

import (
	"fmt"

	pb "elibot-apiserver/serverpb"
	
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type pbserver struct {
}

func (s *pbserver) SayHello(ctx context.Context, in *pb.Req) (*pb.Reply, error) {
		if in.Name == 1 {
			fmt.Println("success")
		}
        return &pb.Reply{Message: "Hello "}, nil
}

func Server() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &pbserver{})
	return s
}