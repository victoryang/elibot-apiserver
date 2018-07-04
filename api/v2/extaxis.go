package v2

import(
	//Log "elibot-apiserver/log"
	"errors"
	"time"
	"strconv"
	"elibot-apiserver/mcserver"

	"golang.org/x/net/context"
	pb "elibot-apiserver/serverpb"
)

var Tag = "grpc:ExtAxis"

const (
	cmd = "goto_extaxispos "
)

type ExtAxis struct {
}

func (e *ExtAxis) GotoExtaxisPos (ctx context.Context, in *pb.Req) (*pb.Reply, error) {
		if McServer == nil {
			return nil, errors.New("mcserver is not available right now")
		}
		
		rCh := make(chan mcserver.Response)
		defer close(rCh)

		c, cancel := context.WithTimeout(ctx, 1*time.Second)
		defer cancel()
		
		command := cmd + strconv.Itoa(int(in.Axis)) + " " + strconv.Itoa(int(in.Num))
		go McServer.Exec(command, Tag, rCh)

		select {
		case <-c.Done():
			return nil, c.Err()
		case r := <- rCh:
			return &pb.Reply{Res: r.Result}, r.Err
		}
}

func NewExtAxisServer() *ExtAxis {
	return &ExtAxis{}
}