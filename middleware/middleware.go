package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func NewCorsHandler(h http.Handler) http.Handler {
	return cors.AllowAll().Handler(h)
}