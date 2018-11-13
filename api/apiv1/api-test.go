package apiv1

import (
	"net/http"
	"context"
	"time"

	"elibot-apiserver/mcserver"
	Log "elibot-apiserver/log"
)

func hello(w http.ResponseWriter, r *http.Request) {
	WriteSuccessResponse(w, "Welcome to elibot")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	WriteSuccessResponse(w, "echo OK")
}