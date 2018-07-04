package v2

import(
	//Log "elibot-apiserver/log"
	"errors"
	"time"
	"elibot-apiserver/mcserver"

	"golang.org/x/net/context"
	pb "elibot-apiserver/serverpb/hello"
)

var test_cmd = "testGo 0 1\n"
/*var cmd = "setZeroEncode 0 1"*/
var from = "grpc:hello"

type helloserver struct {
}

func (s *helloserver) SayHello(ctx context.Context, in *pb.Req) (*pb.Reply, error) {
		if McServer == nil {
			return nil, errors.New("mcserver is not available right now")
		}
		
		rCh := make(chan mcserver.Response)
		defer close(rCh)

		c, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		
		if in.Name == 1 {
			go McServer.Exec(test_cmd, from, rCh)
		} else {
			return &pb.Reply{Message: "test fail"}, errors.New("request name not 1")
		}

		select {
		case <-c.Done():
			return nil, c.Err()
		case r := <- rCh:
			return &pb.Reply{Message: r.Result}, r.Err
		}
}

func NewHelloServer() *helloserver {
	return &helloserver{}
}