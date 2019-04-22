package websocket

import (
	"net/http"
	"time"
	"context"

	Log "elibot-apiserver/log"
	"elibot-apiserver/config"
)

type WsServer struct {
    httpServer      *http.Server
    hub         *Hub
}

var Server *WsServer

// A Graceful shutdown for server
func Shutdown() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    // Doesn't block if no connections, but will otherwise wait
    // until the timeout deadline.
    Server.httpServer.Shutdown(ctx)
    Log.Print("websocket server shuting down...")
}

func PushBytes(msg []byte) {
    Log.Debug("Push messages to all client: ", msg)
    Server.hub.Broadcast(msg)
}

func NotificationRegister(f func([]byte)) {
    Server.hub.NotificationRegister(f)
}

func StartServer(c *config.WebsocketEntryPoint) {
    Server = new(WsServer)

    Server.hub = NewHub()
    go Server.hub.Run()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                ServeHome("assets/index.html",w,r)
        })
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
                ServeWs(Server.hub, w, r)
        })
    Server.httpServer = &http.Server {
        Addr:       c.ListenAddress,
        ReadTimeout:    10 * time.Minute,
        WriteTimeout:   10 * time.Minute,
        MaxHeaderBytes: 1 << 20,
    }

    Log.Print("websocket server listening...")
    go Server.httpServer.ListenAndServe()
}
