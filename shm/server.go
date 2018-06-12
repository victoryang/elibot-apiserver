package shm

// #cgo CFLAGS: -I/root/mcserver/include/
// #cgo LDFLAGS: -lshare
// #include<stdlib.h>
// #include<sharedmemory.h>
import "C"
import(
	"context"
	"time"
	"fmt"
	"elibot-apiserver/api"
	Log "elibot-apiserver/log"
)

var hit chan []byte
const (
	BUFSIZE = 1024
)

type ShmServer struct {
	Wss 		*api.WsServer
	Hit 		chan []byte
	Ctx 		context.Context
	Cancel 		context.CancelFunc
}

func (s *ShmServer) Watcher() {
	Log.Info("Shared memory server started...")
	watchTicker := time.NewTicker(watchPeriod)
	defer func() {
		watchTicker.Stop()
	}()

	DONE:
	for {
		select {
		case <-s.Ctx.Done():
			break DONE
		case res := <-modified:
			s.Wss.Hub.PushMsg(res)
		}
	}
}

func (s *ShmServer) Shutdown() {
	s.Cancel()
}

func handleMsg(msg []byte) {
	fmt.Println(string(msg[:]))
}

func (s *ShmServer) StartToWatch() {
	go worker(s.Ctx, s.Hit)
	go s.Watcher()
}

func NewServer(server *api.WsServer) *ShmServer{
	s := new(ShmServer)
	ctx, cancel := context.WithCancel(context.Background())
	s = &ShmServer{
		Wss: 	server,
		Hit:	make(chan []byte, BUFSIZE),
		Ctx: 	ctx,
		Cancel: cancel,
	}
	C.init_shm()
	initWatchFuncs()
	server.Hub.NotificationRegister(handleMsg)
	return s
}