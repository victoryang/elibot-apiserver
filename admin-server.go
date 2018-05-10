package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"

	"github.com/gorilla/mux"
)

type handlerFunc func(http.Handler) http.Handler
var handlerFns = []handlerFunc{
//	SetJwtMiddlewareHandler,
}

func RegisterHandlers (mux.Router, handlerFns ...handlerFunc) http.Handler {
	var f http.Handler
	f =mux
	for _, hFn := range handlerFns {
		f = hFn(f)
	}
	return f
}

func configureAdminHandler() http.Handler {
	r := mux.NewRouter()
	apiRouter := r.NewRoute().PathPrefix("/").Subrouter()
	admin := apiRouter.PathPrefix("/admin").Subrouter()

	/*TODO: get some status of controller back to admin*/
	/*admin.Methods("GET").Path("/usage").HandlerFunc(SetJwtMiddlewareFunc(getUsage))*/

	apiRouter.Path("/debug/cmdline").HandlerFunc(pprof.Cmdline)
	apiRouter.Path("/debug/profile").HandlerFunc(pprof.Profile)
	apiRouter.Path("/debug/symbol").HandlerFunc(pprof.Symbol)
	apiRouter.Path("/debug/trace").HandlerFunc(pprof.Trace)
	apiRouter.PathPrefix("/debug/pprof/").HandlerFunc(pprof.Index)

	return RegisterHandlers(r, handlerFns...)
}

func startAdminServer() {
	host := "127.0.0.1"
	port := "9090"
	serverAddress = net.JoinHostPort(host, port)
	adminServer := &http.Server{
		Addr: 			serverAddress,
		// Adding timeout of 10 minutes for unresponsive client connections.
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		Handler:        configureAdminHandler(),
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// Configure TLS if certs are available.
		err = adminServer.ListenAndServe()
		if err!= nil {
			fmt.Println ("API server error.")
		}
		fmt.Println ("API server running...")
	}()
}