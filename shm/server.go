package shm

import(
	"context"
	"fmt"
	"elibot-apiserver/websocket"
	Log "elibot-apiserver/log"
)

var hit chan []byte
const (
	BufferSize = 1024
)

type ShmServer struct {
	Wss 		*websocket.WsServer
	Hit 		chan []byte
	Ctx 		context.Context
	Cancel 		context.CancelFunc
}

func sender(ctx context.Context, ws *websocket.WsServer, hit chan []byte) {
	Log.Print("Shared memory server started...")

	DONE:
	for {
		select {
		case <-ctx.Done():
			break DONE
		case res := <-hit:
			if ws != nil {
				ws.PushBytes(res)
			}
		}
	}
}

func (s *ShmServer) Shutdown() {
	s.Cancel()
	Log.Print("shared memory server shutting down")
}

func handleMsg(msg []byte) {
	fmt.Println(string(msg[:]))
}

func (s *ShmServer) StartToWatch() {
	go worker(s.Ctx, s.Hit)
	go sender(s.Ctx, s.Wss, s.Hit)
}

func NewServer(server *websocket.WsServer) (*ShmServer, error){
	s := new(ShmServer)
	ctx, cancel := context.WithCancel(context.Background())
	s = &ShmServer{
		Wss: 	server,
		Hit:	make(chan []byte, BufferSize),
		Ctx: 	ctx,
		Cancel: cancel,
	}
	err := initWorkerResource()
	if err != nil {
		return nil, err
	}
	initWatchFuncs()
	server.NotificationRegister(handleMsg)
	return s, nil
}