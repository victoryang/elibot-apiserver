package api

import (
	"crypto/tls"
	//"crypto/x509"
	"net/http"
	"time"
	"context"

	Log "elibot-apiserver/log"
	"elibot-apiserver/config"
	"elibot-apiserver/api/v3"
)

const (
	ListenAddress = ":9060"
)

type WsTLSServer struct {
	httpServer 		*http.Server
	hub 			*v3.Hub
}

func (s *WsTLSServer) Run() {
	Log.Debug("websocket server listening...")
	go s.httpServer.ListenAndServeTLS("yourdomain.com.crt", "yourdomain.com.key")
}

// A Graceful shutdown for server
func (s *WsTLSServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    // Doesn't block if no connections, but will otherwise wait
    // until the timeout deadline.
    s.httpServer.Shutdown(ctx)
    Log.Debug("websocket server shuting down...")
}

func (s *WsTLSServer) PushBytes(msg []byte) {
	Log.Debug("Push messages to all client: ", msg)
	s.hub.Broadcast(msg)
}

func NewWssServer(c *config.WebsocketEntryPoint) *WsTLSServer {
	s := new(WsTLSServer)

	s.hub = v3.NewHub()
    go s.hub.Run()
	http.HandleFunc("/", v3.ServeHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
                v3.ServeWs(s.hub, w, r)
    })

    tlsconfig := &tls.Config{}
	// ensure http2 enabled
	tlsconfig.NextProtos = []string{"h2", "http/1.1"}
	s.httpServer = &http.Server {
		Addr:			ListenAddress,
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		MaxHeaderBytes: 1 << 20,
		TLSConfig:		tlsconfig,
	}
	return s
}