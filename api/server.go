package api

import (
	"net"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

type ServerConfig struct {
	addr	string
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

// configureServer handler returns final handler for the http server.
func configServerHandler(c *ServerConfig) http.Handler {
	// Initialize router.
	r := mux.NewRouter().SkipClean(true)
	
	// Register all routers.
	RegisterAPIRouter(r)
	// Add new routers here.

	// List of some generic handlers which are applied for all
	// incoming requests.
	var handlerFns = []api.HandlerFunc{
		// Limits the number of concurrent http requests.
		api.SetCommonHeaderHandler,
		// CORS setting for all browser API requests.
		//api.SetCorsHandler,
		// Validates all incoming URL resources, for invalid/unsupported
		// resources client receives a HTTP error.
		//api.SetIgnoreResourcesHandler,
		// Auth handler verifies incoming authorization headers and
		// routes them accordingly. Client receives a HTTP error for
		// invalid/unsupported signatures.
		//api.SetAuthHandler,
		// Add new handlers here.

		//api.SetLogHandler,
	}

	// Register rest of the handlers.
	return RegisterHandlers(r, handlerFns...)
}

func NewApiServer(c *ServerConfig) *Server {
	s := new(Server)
	s.entryPoint.httpServer = &http.Server {
		Addr:	c.addr,
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		Handler:        configServerHandler(c),
		MaxHeaderBytes: 1 << 20,
	}
	return s
}