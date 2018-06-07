package api

import (
	"net/http"
	"time"
	"context"

	Log "elibot-apiserver/log"
	
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

type WsServer struct {
	HttpServer 		*http.Server
}

func (s *WsServer) Run() {
	Log.Print("ws server listening...")
	go s.HttpServer.ListenAndServe()
}

// A Graceful shutdown for server
func (s *WsServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    // Doesn't block if no connections, but will otherwise wait
    // until the timeout deadline.
    s.HttpServer.Shutdown(ctx)
    Log.Print("server shuting down...")
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        Log.Error(err)
        return
    }
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "connect to websocket server")
}

func NewWsServer() *WsServer {
	s := new(WsServer)
	
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	s.HttpServer = &http.Server {
		Addr:			"127.0.0.1:9050",
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}