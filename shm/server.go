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
	Ctx 		context.Context
	Cancel 		context.CancelFunc
}

func (s *ShmServer) StartToWatch(ctx context.Context, modified chan []byte) {
	Log.Info("Shared memory server started...")
	go func(ctx context.Context) {
		watchTicker := time.NewTicker(watchPeriod)
		defer func() {
			watchTicker.Stop()
		}()

		DONE:
		for {
			select {
			case <-ctx.Done():
				break DONE
			case res := <-modified:
				s.Wss.Hub.PushMsg(res)
			}
		}
	}(ctx)
}

func (s *ShmServer) Shutdown() {
	s.Cancel()
}

func handleMsg(msg []byte) {
	fmt.Println(string(msg[:]))
}

func (s *ShmServer)initServer() {
	C.init_shm()
	hit = make(chan []byte, BUFSIZE)
	initWatchFuncs()
	s.Wss.Hub.NotificationRegister(handleMsg)
	go worker(s.Ctx, hit)
}

func NewServer(server *api.WsServer) *ShmServer{
	s := new(ShmServer)
	ctx, cancel = context.WithCancel(context.Background())
	s = &ShmServer{
		Wss: 	server,
		Ctx: 	ctx,
		Cancel: cancel,
	}
	s.initServer()
	return s
}