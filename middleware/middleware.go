package middleware

import (
	"github.com/rs/cors"
)

func NewCorsHandler() *cors.Cors {
	return cors.AllowAll()
}