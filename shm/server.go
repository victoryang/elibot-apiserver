package shm

// #cgo CFLAGS: -I/root/mcserver/include/
// #cgo LDFLAGS: -lshare
// #include<stdlib.h>
// #include<sharedmemory.h>
import "C"
import(
	"context"
	"time"
	"elibot-apiserver/api"
	Log "elibot-apiserver/log"
)

var hit chan []byte
const (
	BUFSIZE = 256
)

type ShmServer struct {
	Wss 		*api.WsServer
	Ctx 		context.Context
	Cancel 		context.CancelFunc
}

func (s *ShmServer) StartToWatch(ctx context.Context, modified chan []byte) {
	Log.Info("Shared memory server started...")
	go func(ctx) {
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
				wss.Hub.PushMsg(data)
			}
		}
	}()
}

func (s *ShmServer) Shutdown() {
	s.Cancel()
}

func handleMsg(msg []byte) {
	fmt.Println(string(msg[:]))
}

func init(ctx context.Context) {
	ctx, cancel = context.WithCancel(context.Background())
	C.init_shm()
	hit = make(chan []byte, BUFSIZE)
	initWatchFuncs(hit)
	wss.Hub.NotificationRegister(handleMsg)
	go watcher(ctx, hit)
	go worker(ctx, hit)
}

func NewServer(s *api.WsServer) *ShmServer{
	init()
	
	return &ShmServer{
		Wss: 	s,
		Ctx: 	ctx,
		Cancel: cancel,
	}
}