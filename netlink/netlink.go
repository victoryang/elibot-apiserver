package netlink

import (
	"context"
	"os"
	"encoding/json"
	"syscall"

	Log "elibot-apiserver/log"
	"elibot-apiserver/websocket"
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
		websocket.PushBytes(message)
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