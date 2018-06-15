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
	HttpServer 		*http.Server
	Hub 			*v3.Hub
}

func (s *WsServer) Run() {
	Log.Debug("websocket server listening...")
	go s.HttpServer.ListenAndServe()
}

// A Graceful shutdown for server
func (s *WsServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    // Doesn't block if no connections, but will otherwise wait
    // until the timeout deadline.
    s.HttpServer.Shutdown(ctx)
    Log.Debug("websocket server shuting down...")
}

func NewWsServer(c *config.WebsocketEntryPoint) *WsServer {
	s := new(WsServer)
	
	s.Hub = v3.NewHub()
    go s.Hub.Run()
	http.HandleFunc("/", v3.ServeHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
                v3.ServeWs(s.Hub, w, r)
    })
	s.HttpServer = &http.Server {
		Addr:			c.ListenAddress,
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}