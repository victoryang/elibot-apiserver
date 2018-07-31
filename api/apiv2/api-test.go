package apiv2

import (
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	WriteSuccessResponse(w, "Welcome to elibot\n")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	WriteSuccessResponse(w, "echo OK\n")
}