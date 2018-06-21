package v2

import(
	//Log "elibot-apiserver/log"
	"errors"
	"elibot-apiserver/mcserver"

	"golang.org/x/net/context"
	pb "elibot-apiserver/serverpb"
)

var cmd = "testGo 0 1\n"
/*var cmd = "setZeroEncode 0 1"*/
var from = "grpc:hello"

type helloserver struct {
}

func (s *helloserver) SayHello(ctx context.Context, in *pb.Req) (*pb.Reply, error) {
		var res string = "test Go fail"
		var err error
		var mcs *mcserver.MCserver 
		if mcs = mcserver.GetMcServer(); mcs == nil {
			return nil, errors.New("mcserver is not available right now")
		}
		
		rCh := make(chan mcserver.Response)
		if in.Name == 1 {
			mcs.Exec(ctx, cmd, from, rCh)
			r := <- rCh
			res = r.Result
			err = r.Err
		}
        return &pb.Reply{Message: res}, err
}

func NewHelloServer() *helloserver {
	return &helloserver{}
}