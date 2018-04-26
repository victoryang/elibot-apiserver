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

func testv1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "echo")
	return
}

func RegisterAPIRouter(r *mux.Router) *mux.Router {
	r.HandleFunc("/", hello).Methods("GET")
	r.HandleFunc("/api/v1/testv1", testv1).Methods("GET")
	return r
}