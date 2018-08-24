package mcserver

import (
	"errors"
	"net"
	"context"

	Log "elibot-apiserver/log"
)

type Request struct {
	Command 		string
	From			string
	RespCh 			chan Response
}

type Response struct {
	Result 			string
	Err 			error
}

func getConnFromPool() (interface{}, Response) {
	Log.Debug("MCServer get a connection from pool...")
	if Mcs == nil {
		return nil, Response{Result: "", Err: errors.New("MCServer Error")}
	}

	conn, err := Mcs.ConnPool.Get()
	if err!=nil {
		return nil, Response{Result: "", Err: err}
	}
	return conn, Response{}
}

func execute(ctx context.Context, ch chan Response, cmd string) {
	Log.Debug("MCServer executing new command...")
	conn, resp := getConnFromPool()
	if conn == nil {
		SafeSendResponseToChannel(ch, resp)
		return
	}
	defer func() {
		Mcs.ConnPool.Put(conn)
	}()

	select {
	case <-ctx.Done():
		Log.Debug("job cancelled")
		SafeSendResponseToChannel(ch, Response{Result: "", Err: errors.New("MCserver job cancelled")})
		return
	default:
		res, err := HandleCommand(conn.(net.Conn), cmd)
		SafeSendResponseToChannel(ch, Response{Result: res, Err: err})
	}
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case req := <- Mcs.WorkChan:
			go execute(ctx, req.RespCh, req.Command)
		}
	}
}