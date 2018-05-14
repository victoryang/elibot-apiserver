package api

import (
	"net/http"
	"time"

	Log "elibot-apiserver/log"

	"elibot-apiserver/middleware"
	"elibot-apiserver/middleware/accesslog"
	"elibot-apiserver/config"
	
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type AccessLogMiddleware struct {
	File 		string
	Logger 		*accesslog.Logger
}

type ServerEntryPoint struct {
	httpServer       *http.Server
}

type Server struct {
	EntryPoint      	ServerEntryPoint
	AccessLog			AccessLogMiddleware
}

func (s *Server) Run() {
	Log.Print("server listening...")
	go s.EntryPoint.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() {
	Log.Print("server shuting down...")
	if s.AccessLog.Logger!=nil {
		s.AccessLog.Logger.Close()
	}
	s.EntryPoint.httpServer.Close()
}

// configureServer handler returns final handler for the http server.
func (s *Server) configServerHandler() http.Handler {
	// Initialize router.
	r := mux.NewRouter().SkipClean(true)

	n := negroni.New(negroni.NewRecovery())

	s.AccessLog.Logger = middleware.AddAccesslog(n, s.AccessLog.File)
	if s.AccessLog.Logger!=nil {
		n.Use(s.AccessLog.Logger)
	}
	
	n.UseHandler(RegisterAPIRouter(r))
	
	// Register all routers.
	return n
}

func NewApiServer(c *config.Config) *Server {
	s := new(Server)
	
	s.AccessLog.File = c.AccessLogsFile
	s.EntryPoint.httpServer = &http.Server {
		Addr:			c.ListenAddress,
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		Handler:        s.configServerHandler(),
		MaxHeaderBytes: 1 << 20,
	}
	return s
}