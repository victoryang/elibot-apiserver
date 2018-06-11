package v2

import(
	//Log "elibot-apiserver/log"

	"elibot-apiserver/mcserver"

	"golang.org/x/net/context"
	pb "elibot-apiserver/serverpb"
)

var test = "testGo 0 1"
var end = "\n"

type helloserver struct {
}

func (s *helloserver) SayHello(ctx context.Context, in *pb.Req) (*pb.Reply, error) {
		var res string = "test Go fail"
		var err error
		
		ctx_m, _ := context.WithTimeout(ctx, 1000*time.Millisecond)
		if in.Name == 1 {
			cmd := test+end
			res, err = mcserver.OnCommandRecived(ctx_m, cmd)
		}
        return &pb.Reply{Message: res}, err
}

func NewHelloServer() *helloserver {
	return &helloserver{}
}