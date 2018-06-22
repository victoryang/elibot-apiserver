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
		var mcs *mcserver.MCserver 
		if mcs = mcserver.GetMcServer(); mcs == nil {
			return nil, errors.New("mcserver is not available right now")
		}
		
		select {
		case <-ctx.Done():
			return nil, errors.New("test Go cancelled")
		default:
			rCh := make(chan mcserver.Response)
			defer close(rCh)

			var r mcserver.Response
			if in.Name == 1 {
				go mcs.Exec(cmd, from, rCh)
				r = <- rCh
			}
			return &pb.Reply{Message: r.Result}, r.Err
		}
}

func NewHelloServer() *helloserver {
	return &helloserver{}
}