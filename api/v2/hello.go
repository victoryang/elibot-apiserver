package v2

import(
	//Log "elibot-apiserver/log"

	"elibot-apiserver/mcserver"

	"golang.org/x/net/context"
	pb "elibot-apiserver/serverpb"
)

type helloserver struct {
}

func (s *helloserver) SayHello(ctx context.Context, in *pb.Req) (*pb.Reply, error) {
		var res string = "test Go fail"
		if in.Name == 1 {
			res, _ = mcserver.OnCommandRecived()
		}
        return &pb.Reply{Message: res}, nil
}

func NewHelloServer() *helloserver {
	return &helloserver{}
}