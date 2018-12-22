package netlink

import (
	"context"
	"os"
	"encoding/json"
	"syscall"

	Log "elibot-apiserver/log"
	"elibot-apiserver/paramserver"
)

const (
	serviceMethod = "push_message_to_network"
)

var ctx, cancel = context.WithCancel(context.Background())

type WsResponse struct {
	UdiskEvent		UEvent 			`json:"udiskEvent"`
}

type NLServer struct {
	fd		int
}

func (nl *NLServer) Close() error {
	cancel()
	return syscall.Close(nl.fd)
}

func handleMsg(msg []byte) {
	u := parseUEvent(msg)
	if u == nil || u.Env[SUBSYSTEM] != "usb" {
		return
	}

	message, err := json.Marshal(WsResponse{UdiskEvent: *u})
	if err!=nil {
		Log.Error("Could not marshal to json ", err)
	} else {
		var reply bool
		params := make(map[string]interface{})
		params["message"] = string(message)

	    err := paramserver.SendToParamServerWithJsonRpc(serviceMethod, params, &reply)
	    if err!=nil {
	        Log.Error("Could not call rpc request to param server: ", err)
	        return
	    }

	    Log.Debug("reply is ", reply)
	}
	return
}

func (nl *NLServer) Start() {
	go func() {	
		for {
			select {
			case <-ctx.Done():
				return
			default:
				buf := make([]byte, 256)
				n, err := syscall.Read(nl.fd, buf)
				if err!=nil {
					Log.Error("Read err, return...")
					return
				}

				go handleMsg(buf[:n])
			}
		}
	}()
}

func NewNetlinkServer() *NLServer {
	fd, err := syscall.Socket(
		syscall.AF_NETLINK,
		syscall.SOCK_RAW,
		syscall.NETLINK_KOBJECT_UEVENT,
	)
	if err!=nil {
		Log.Error("Could not create socket for netlink: ", err)
		return nil
	}

	addr := syscall.SockaddrNetlink{
		Family: syscall.AF_NETLINK,
		Pid:	uint32(os.Getpid()),
		Groups:	1,
	}
	if err := syscall.Bind(fd, &addr); err!=nil {
		Log.Error("Could not bind socket for netlink: ", err)
		syscall.Close(fd)
		return nil
	}

	return &NLServer{fd}
}