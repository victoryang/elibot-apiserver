package v2

import(
	//Log "elibot-apiserver/log"

	"elibot-apiserver/mcserver"

	"golang.org/x/net/context"
	pb "elibot-apiserver/serverpb"
)

var cmd = "testGo 0 1\n"
var from = "grpc:hello"

type helloserver struct {
}

func (s *helloserver) SayHello(ctx context.Context, in *pb.Req) (*pb.Reply, error) {
		var res string = "test Go fail"
		var err error
		
		if in.Name == 1 {
			resp := make(chan mcserver.Response)
			mcserver.WorkChan<-{Command: cmd, From:from, Resp:resp}
			r<-resp
			res = r.Result
			err = r.Err
		}
        return &pb.Reply{Message: res}, err
}

func NewHelloServer() *helloserver {
	return &helloserver{}
}