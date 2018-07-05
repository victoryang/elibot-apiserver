package apiv1

import (
	"net/http"
	"fmt"
)

// WriteSuccessResponse write success headers and response if any.
func WriteSuccessResponse(w http.ResponseWriter, response string) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, response)
}

func WriteInternalServerErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, err.Error())
}