package apiv1

import (
	"net/http"
	"strings"

	Log "elibot-apiserver/log"
)

func getSourceIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func getTokenFromHeader(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return ""
	}

	parts := strings.Split(auth, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		Log.Error("Authorization header format must be Bearer {token}")
		return ""
	}

	return parts[1]
}