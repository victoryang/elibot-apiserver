package api

import (
	"net/http"
	"time"
	"context"

	Log "elibot-apiserver/log"
	"elibot-apiserver/config"
	"elibot-apiserver/api/v3"
)

type WsServer struct {
	httpServer 		*http.Server
	hub 			*v3.Hub
}

func (s *WsServer) Run() {
	Log.Print("websocket server listening...")
	go s.httpServer.ListenAndServe()
}

// A Graceful shutdown for server
func (s *WsServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    // Doesn't block if no connections, but will otherwise wait
    // until the timeout deadline.
    s.httpServer.Shutdown(ctx)
    Log.Print("websocket server shuting down...")
}

func (s *WsServer) PushBytes(msg []byte) {
	Log.Debug("Push messages to all client: ", msg)
	s.hub.Broadcast(msg)
}

func (s *WsServer) NotificationRegister(f func([]byte)) {
	s.hub.NotificationRegister(f)
}

func NewWsServer(c *config.WebsocketEntryPoint) *WsServer {
	s := new(WsServer)
	
	s.hub = v3.NewHub()
    go s.hub.Run()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                v3.ServeHome(c.IndexFile, w, r)
    })
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
                v3.ServeWs(s.hub, w, r)
    })
	s.httpServer = &http.Server {
		Addr:			c.ListenAddress,
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}