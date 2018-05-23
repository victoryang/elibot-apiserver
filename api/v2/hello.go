package v2

import(
	Log "elibot-apiserver/log"
	
	"golang.org/x/net/context"
	pb "elibot-apiserver/serverpb"
)

type helloserver struct {
}

func (s *helloserver) SayHello(ctx context.Context, in *pb.Req) (*pb.Reply, error) {
		if in.Name == 1 {
			Log.Println("success")
		}
        return &pb.Reply{Message: "Hello "}, nil
}

func NewHelloServer() *helloserver {
	return &helloserver{}
}