package api

import (
	"net/http"

	"elibot-apiserver/api/v1"
	Log "elibot-apiserver/log"

	"github.com/gorilla/mux"
)

func RegisterAPIRouter(r *mux.Router) http.Handler {
	Log.Debug("Register V1 api...")
	return v1.RegisterV1(r)
}