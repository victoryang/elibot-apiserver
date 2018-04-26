package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "test")
	return
}

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "echo")
	return
}

func RegisterAPIRouter(r *mux.Router) http.Handler {
	r.HandleFunc("/", hello).Methods("GET")
	r.HandleFunc("/api/v1/test", testv1).Methods("GET")
	return r
}