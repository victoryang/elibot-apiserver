package api

import (
	"github.com/gorilla/mux"
)

func RegisterAPIRouter(r *mux.Router) *mux.Router {
	r.HandleFunc("/",
	r.HandleFunc("/api/v1/testv1", testv1).Methods("GET")
	return r
}