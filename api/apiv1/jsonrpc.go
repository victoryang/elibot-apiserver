package apiv1

import (
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	Log "elibot-apiserver/log"
)

var JsonRpcClient *rpc.Client
var address string = "127.0.0.1:8055"

func NewJsonRpcClient() {
	var err error
	if JsonRpcClient, err = jsonrpc.Dial("tcp", address); err!=nil {
		Log.Error("Could not dial to mcserver")
	}
	return
}

func SendToMCServerWithJsonRpc(w http.ResponseWriter, serviceMethod string, args interface{}) {
	var reply interface{}
	if err:=JsonRpcClient.Call(serviceMethod, args, reply); err!=nil {
		Log.Error("Could not call request by rpc")
		WriteInternalServerErrorResponse(w, ERRMCSEVERNOTAVAILABLE)
	}

	WriteSuccessResponse(w, reply)
}