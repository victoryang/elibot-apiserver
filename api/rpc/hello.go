package rpc

import(
	//Log "elibot-apiserver/log"
	"errors"
	"time"

	"golang.org/x/net/context"
	pb "elibot-apiserver/serverpb/hello"
)

var test_cmd = "testGo 0 1\n"
/*var cmd = "setZeroEncode 0 1"*/
var from = "grpc:hello"

type helloserver struct {
}

func (s *helloserver) SayHello(ctx context.Context, in *pb.Req) (*pb.Reply, error) {	
		rCh := make(chan string)
		defer close(rCh)

		c, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		
		if in.Name == 1 {
			go SendToMCServer(command, nil, rCh)
		} else {
			return &pb.Reply{Message: "test fail"}, errors.New("request name not 1")
		}

		select {
		case <-c.Done():
			return nil, c.Err()
		case r := <- rCh:
			return &pb.Reply{Message: r}, nil
		}
}

func NewHelloServer() *helloserver {
	return &helloserver{}
}