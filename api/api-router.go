package api

import (
	"github.com/gorilla/mux"
)


func RegisterAPIRouter(r *mux.Router) *mux.Router {
	r.HandlerFunc("/api/v1/testv1", testv1).Method("GET")
	return r
}

func RegisterHandlers() (r *mux.Router){
	return
}