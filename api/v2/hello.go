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
		var err error
		
		go func() {
			if in.Name == 1 {
				res, err = mcserver.OnCommandRecived()
			}
			<-ctx.Done()
		}()
        return &pb.Reply{Message: res}, err
}

func NewHelloServer() *helloserver {
	return &helloserver{}
}