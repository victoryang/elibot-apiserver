package paramserver

import (
    "context"
    "net"
    "time"
    "sync"
    "elibot-apiserver/jsonrpc2"

    Log "elibot-apiserver/log"
)

var ParamServerRpcClient *jsonrpc2.Conn
var address string = "127.0.0.1:9030"
var ctx_rpc, cancel_rpc = context.WithCancel(context.Background())
var JsonRpcTimeOut = 5 * time.Second
var mux_rpc sync.Mutex

type handler struct {
}

func (h *handler)Handle(ctx context.Context, c *jsonrpc2.Conn, req *jsonrpc2.Request) {
    Log.Debug("Receive request from server")
    return
}

func CloseRpcClient() {
    cancel_rpc()
    ParamServerRpcClient.Close()
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
    ParamServerRpcClient = jsonrpc2.NewConn(ctx_rpc, stream, new(handler))
    return nil
}

func NewRpcClient() error {
    conn, err := net.Dial("tcp", address)
    if err != nil {
        Log.Error("Dial: ", err)
        return err
    }
    stream := jsonrpc2.NewBufferedStream(conn, jsonrpc2.VarintObjectCodec{})
    ParamServerRpcClient = jsonrpc2.NewConn(ctx_rpc, stream, new(handler))
    return nil
}

func SendToParamServerWithJsonRpc(serviceMethod string, params interface{}, reply interface{}) error {
    ctx, cancel := context.WithTimeout(ctx_rpc, JsonRpcTimeOut)
    defer cancel()

    Log.Debug("send to param server, serviceMethod: ", serviceMethod, " params: ", params)

    mux_rpc.Lock()
    defer mux_rpc.Unlock()
    if ParamServerRpcClient.IsClosed() {
        // Current connection closed, reconnect 3 times
        if err := reconnect(); err!=nil {
            Log.Error("Could not reconnect to param server")
            return err
        }
    }

    if err:=ParamServerRpcClient.Call(ctx, serviceMethod, params, reply); err!=nil {
        Log.Error("Could not call request by rpc: ", err)
        return err
    }

    return nil
}