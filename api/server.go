package api

import (
	"net/http"
	"time"
	"fmt"
	"elibot-apiserver/config"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type ServerEntryPoint struct {
	httpServer       *http.Server
}

type Server struct {
	EntryPoint      ServerEntryPoint
}

func (s *Server) Run() {
	fmt.Println("server listening...")
	go s.EntryPoint.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() {
	fmt.Println("server shuting down...")
	s.EntryPoint.httpServer.Close()
}

// configureServer handler returns final handler for the http server.
func configServerHandler() http.Handler {
	// Initialize router.
	r := mux.NewRouter().SkipClean(true)

	n := negroni.Classic()
	n.UseHandler(RegisterAPIRouter(r))
	
	// Register all routers.
	return n
}

func NewApiServer(c config.GlobalConfiguration) *Server {
	s := new(Server)

	s.EntryPoint.httpServer = &http.Server {
		Addr:			c.ListenAddress,
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		Handler:        configServerHandler(),
		MaxHeaderBytes: 1 << 20,
	}
	return s
}