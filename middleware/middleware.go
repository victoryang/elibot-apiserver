package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func NewCorsHandler() http.Handler {
	return cors.AllowAll().Handler()
}