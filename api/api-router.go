package api

import (
	"net/http"
	"fmt"

	"elibot-apiserver/api/v1"

	"github.com/gorilla/mux"
)

func RegisterAPIRouter(r *mux.Router) http.Handler {
	fmt.Println("Register V1 api...")
	return v1.RegisterV1(r)
}