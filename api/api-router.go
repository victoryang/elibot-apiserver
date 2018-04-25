package api

import (
	"github.com/gorilla/mux"
)


func RegisterAPIRouter(r *mux.Router) (r *mux.Router) {
	r.HandlerFunc("/api/v1/test", testv1).Method("GET")
}

func RegisterHandlers() (r *mux.Router){
	return
}