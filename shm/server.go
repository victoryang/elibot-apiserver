package shm

import(
	"context"
	"time"
	"fmt"
	"elibot-apiserver/api"
	Log "elibot-apiserver/log"
)

var hit chan []byte
const (
	BufferSize = 1024
)

type ShmServer struct {
	Wss 		*api.WsServer
	Hit 		chan []byte
	Ctx 		context.Context
	Cancel 		context.CancelFunc
}

func sender(ctx context.Context, ws *api.WsServer, hit chan []byte) {
	Log.Info("Shared memory server started...")
	watchTicker := time.NewTicker(watchPeriod)
	defer func() {
		watchTicker.Stop()
	}()

	DONE:
	for {
		select {
		case <-ctx.Done():
			break DONE
		case res := <-hit:
			ws.Hub.PushMsg(res)
		}
	}
}

func (s *ShmServer) Shutdown() {
	s.Cancel()
	Log.Info("sharedmemory server shutting down")
}

func handleMsg(msg []byte) {
	fmt.Println(string(msg[:]))
}

func (s *ShmServer) StartToWatch() {
	go worker(s.Ctx, s.Hit)
	go sender(s.Ctx, s.Wss, s.Hit)
}

func NewServer(server *api.WsServer) *ShmServer{
	s := new(ShmServer)
	ctx, cancel := context.WithCancel(context.Background())
	s = &ShmServer{
		Wss: 	server,
		Hit:	make(chan []byte, BufferSize),
		Ctx: 	ctx,
		Cancel: cancel,
	}
	initWorkerResource()
	initWatchFuncs()
	server.Hub.NotificationRegister(handleMsg)
	return s
}