package apiv1

import (
	"context"
	"net"
	"net/http"
	"time"
	"sync"
	"elibot-apiserver/jsonrpc2"

	Log "elibot-apiserver/log"
)

var JsonRpcClient *jsonrpc2.Conn
var address string = "127.0.0.1:8055"
var ctx_rpc, cancel_rpc = context.WithCancel(context.Background())
var JsonRpcTimeOut = 5 * time.Second
var mux_rpc sync.Mutex

type handler struct {
}

func (h *handler)Handle(ctx context.Context, c *jsonrpc2.Conn, req *jsonrpc2.Request) {
	Log.Debug("Receive request from server")
	return
}

func CloseJsonRpcClient() {
	cancel_rpc()
	JsonRpcClient.Close()
}

func reconnect() error {
	var maxRetryTimes = 3
	var conn net.Conn
	var err error
	for maxRetryTimes > 0 {
		conn, err = net.Dial("tcp", address)
		if err == nil {
			break
		} else {
			maxRetryTimes--
		}
	}

	if err != nil {
		return err
	}

	stream := jsonrpc2.NewBufferedStream(conn, jsonrpc2.VarintObjectCodec{})
	JsonRpcClient = jsonrpc2.NewConn(ctx_rpc, stream, new(handler))
	return nil
}

func NewJsonRpcClient() error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		Log.Error("Dial: ", err)
		return err
	}
	stream := jsonrpc2.NewBufferedStream(conn, jsonrpc2.VarintObjectCodec{})
	JsonRpcClient = jsonrpc2.NewConn(ctx_rpc, stream, new(handler))
	return nil
}

func ConcatParams(args ...string) []string {
	var params []string
	for _,v := range args {
		if v == "" {
			continue
		}
		params = append(params, v)
	}

	return params
}

func SendToMCServerWithJsonRpc(w http.ResponseWriter, serviceMethod string, params []string) {
	var reply bool

	ctx, cancel := context.WithTimeout(ctx_rpc, JsonRpcTimeOut)
	defer cancel()

	mux_rpc.Lock()
	if JsonRpcClient.GetClosingStatus() {
		// Current connection closed, reconnect 3 times
		if err := reconnect(); err!=nil {
			Log.Error("Could not reconnect to mcserver")
			WriteInternalServerErrorResponse(w, ERRMCSEVERNOTAVAILABLE)
			mux_rpc.Unlock()
			return
		}
	}
	mux_rpc.Unlock()

	Log.Debug("send to server, serviceMethod: ", serviceMethod, " params: ", params)
	if len(params) == 0 {
		if err:=JsonRpcClient.Call(ctx, serviceMethod, nil, &reply); err!=nil {
			Log.Error("Could not call request by rpc: ", err)
			WriteInternalServerErrorResponse(w, ERRMCSEVERNOTAVAILABLE)
			return
		}
	} else {
		if err:=JsonRpcClient.Call(ctx, serviceMethod, params, &reply); err!=nil {
			Log.Error("Could not call request by rpc: ", err)
			WriteInternalServerErrorResponse(w, ERRMCSEVERNOTAVAILABLE)
			return
		}
	}

	WriteSuccessResponse(w, reply)
}