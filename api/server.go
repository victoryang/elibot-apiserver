package api

import (
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

type ServerConfig struct {
	Addr	string
}

type ServerEntryPoint struct {
	httpServer       *http.Server
}

type Server struct {
	entryPoint      ServerEntryPoint
}

func (s *Server) Run() {
	s.entryPoint.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() {
	s.entryPoint.httpServer.Shutdown()
}

// configureServer handler returns final handler for the http server.
func configServerHandler(c *ServerConfig) http.Handler {
	// Initialize router.
	r := mux.NewRouter().SkipClean(true)
	
	// Register all routers.
	return RegisterAPIRouter(r)
}

func NewApiServer(c *ServerConfig) *Server {
	s := new(Server)
	s.entryPoint.httpServer = &http.Server {
		Addr:	c.Addr,
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		Handler:        configServerHandler(c),
		MaxHeaderBytes: 1 << 20,
	}
	return s
}