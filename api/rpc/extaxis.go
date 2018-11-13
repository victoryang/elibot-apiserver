package rpc

import(
	//Log "elibot-apiserver/log"
	"time"

	"golang.org/x/net/context"
	pb "elibot-apiserver/serverpb"
)

var Tag = "grpc:ExtAxis"

const (
	cmd = "goto_extaxispos"
)

type ExtAxis struct {
}

func (e *ExtAxis) GotoExtaxisPos (ctx context.Context, in *pb.Req) (*pb.Reply, error) {		
		rCh := make(chan string)
		defer close(rCh)

		c, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		go SendToMCServer(cmd, nil, rCh)

		select {
		case <-c.Done():
			return nil, c.Err()
		case r := <- rCh:
			return &pb.Reply{Res: r}, nil
		}
}

func NewExtAxisServer() *ExtAxis {
	return &ExtAxis{}
}