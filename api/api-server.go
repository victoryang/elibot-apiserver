package api

import (
	"net/http"
	"time"

	Log "elibot-apiserver/log"
	"elibot-apiserver/middleware/accesslog"
	"elibot-apiserver/config"
	
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

var accesslogfile string

type ServerEntryPoint struct {
	httpServer       *http.Server
}

type Server struct {
	EntryPoint      ServerEntryPoint
}

func (s *Server) Run() {
	Log.Print("server listening...")
	go s.EntryPoint.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() {
	Log.Print("server shuting down...")
	s.EntryPoint.httpServer.Close()
}

// configureServer handler returns final handler for the http server.
func configServerHandler() http.Handler {
	// Initialize router.
	r := mux.NewRouter().SkipClean(true)

	n := negroni.New(negroni.NewRecovery())
	alogger, err := accesslog.NewAccessLog(accesslogfile)
	if err!=nil {
		Log.Error("Error in opening access log file： ", err)
	} else {
		n.Use(alogger)
	}
	/*static := negroni.NewStatic()*/
	
	n.UseHandler(RegisterAPIRouter(r))
	
	// Register all routers.
	return n
}

func NewApiServer(c *config.GlobalConfiguration) *Server {
	s := new(Server)
	
	accesslogfile = c.AccessLogsFile
	s.EntryPoint.httpServer = &http.Server {
		Addr:			c.ListenAddress,
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		Handler:        configServerHandler(),
		MaxHeaderBytes: 1 << 20,
	}
	return s
}