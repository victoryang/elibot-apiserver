package main 

import(
	"net/http"
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func handleSignals(s *Server) {
	signal.Ignore()
	signalQueue := make(chan os.Signal)
	signal.Notify(signalQueue, syscall.SIGHUP, os.Interrupt)
	for {
		sig := <-signalQueue
		switch sig {
		//TODO:
		//case syscall.SIGHUP:
			//reload config file
		default:
			// stop server
			s.Stop()
			fmt.Println("Admin server stop...")
			os.Exit(0)
		}
	}
}

type Server struct {
	admin 		*http.Server
	monitor		*Monitor
}

const (
	DefaultInterval = 3*time.Second
	DefaultRequestTimeout = 2*time.Second
)

func (m *Monitor) registerbackend(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	url := vars["url"]

	b := NewBackendConfig(Options{Interval: DefaultInterval}, url, DefaultRequestTimeout)

	m.registerCh<-RegisterReq{name, b}
}

func (m *Monitor) unregisterbackend(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	m.unregisterCh<-vars["name"]
}

func (m *Monitor) getBackendHealth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(m.backends[vars["name"]])
}

func (m *Monitor) getBackendMetrics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if res := m.GetStats(vars["name"]); res!=nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, res)
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func (m *Monitor) getBackends(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var res []string
	for key, _ := range m.backends {
		res = append(res, key)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
}

func serveLogin(w http.ResponseWriter, r *http.Request) {
	
}

func configureHandler(m *Monitor) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/login", serveLogin).Methods("GET").Queries("username", "{username}","pass", "{pass}")

	monitorRouter := r.PathPrefix("/monitor").Subrouter()
	monitorRouter.HandleFunc("/registerbackend/{name}/{url}", m.registerbackend).Methods("POST")
	monitorRouter.HandleFunc("/unregisterbackend/{name}", m.unregisterbackend).Methods("DELETE")
	monitorRouter.HandleFunc("/{name}/health", m.getBackendHealth).Methods("GET")
	monitorRouter.HandleFunc("/{name}/metrics", m.getBackendMetrics).Methods("GET")
	monitorRouter.HandleFunc("/", m.getBackends).Methods("GET")

	apiRouter := r.PathPrefix("/").Subrouter()
	apiRouter.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static/"))))

	return r
}

func startAdminServer() *Server{
	s := new(Server)
	s.monitor = NewMonitor()
	s.admin = &http.Server{
		Addr: 			":9600",
		// Adding timeout of 10 minutes for unresponsive client connections.
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		Handler:        configureHandler(s.monitor),
		MaxHeaderBytes: 1 << 20,
	}

	// Configure TLS if certs are available.
	go func() {
		fmt.Println("Admin server running...")
		s.admin.ListenAndServe()
	}()

	return s
}

func (s *Server) Stop() {
	s.admin.Close()
	s.monitor.StopMonitor()
}

func main() {
	s := startAdminServer()
	handleSignals(s)
}