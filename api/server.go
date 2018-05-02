package api

import (
	"net/http"
	"time"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
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
	fmt.Println("server listening...")
	go s.entryPoint.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() {
	fmt.Println("server shuting down...")
	s.entryPoint.httpServer.Close()
}

// configureServer handler returns final handler for the http server.
func configServerHandler(c *ServerConfig) http.Handler {
	// Initialize router.
	r := mux.NewRouter().SkipClean(true)

	n := negroni.Classic()
	n.UseHandler(RegisterAPIRouter(r))
	
	// Register all routers.
	return n
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