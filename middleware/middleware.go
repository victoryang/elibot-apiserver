package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func NewCorsHandler() *cors.Cors {
	return cors.AllowAll()
}