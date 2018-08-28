package apiv1

import (
	"context"
	"net"
	"net/http"
	"time"
	"elibot-apiserver/jsonrpc2"

	Log "elibot-apiserver/log"
)

var JsonRpcClient *rpc.Client
var address string = "127.0.0.1:8055"
var ctx_base, _ = context.WithCancel(context.Background())
var JsonRpcTimeOut = 5

type handler struct {
}

func (h *handler)Handle(ctx context.Context, c *jsonrpc2.Conn, req *jsonrpc2.Request) {
	fmt.Println("Receive request from server")
	return
}

func NewJsonRpcClient() error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		Log.Error("Dial: ", err)
		return err
	}
	stream := NewBufferedStream(conn, jsonrpc2.VarintObjectCodec{})
	JsonRpcClient = jsonrpc2.NewConn(ctx, stream, new(handler))
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

func SendToMCServerWithJsonRpc(w http.ResponseWriter, serviceMethod string, params interface{}) {
	var reply bool

	ctx, cancel := context.WithTimeout(ctx_base, JsonRpcTimeOut * time.Second)
	defer cancel()

	if err:=JsonRpcClient.Call(ctx, serviceMethod, params, &reply); err!=nil {
		Log.Error("Could not call request by rpc: ", err)
		WriteInternalServerErrorResponse(w, ERRMCSEVERNOTAVAILABLE)
	}

	WriteSuccessResponse(w, reply)
}