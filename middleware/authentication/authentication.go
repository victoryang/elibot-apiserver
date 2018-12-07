package authentication

import (
	"net/http"
	//"context"
	"strings"
	"elibot-apiserver/auth"

	Log "elibot-apiserver/log"

    "github.com/gorilla/mux"

)

type Authenticator struct {
	r 		*mux.Router
}

func NewAuthMiddleware(router *mux.Router) *Authenticator {
	a := new(Authenticator)
    a.r = router
    return a
}

func getSourceIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func getTokenFromHeader(r *http.Request, token *string) bool {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return false
	}

	parts := strings.Split(auth, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		Log.Error("Authorization header format must be Bearer {token}")
		return false
	}

	*token = parts[1]
	return true
}

func (a *Authenticator) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ip := getSourceIP(r)

	//ctx := context.WithValue(r.Context(), "ip", ip)

	if auth.CheckSession(ip) {
		/*var token string
		if getTokenFromHeader(r, &token) {		
			ctx = context.WithValue(ctx, "ip", ip)
		}*/
	}
	var m mux.RouteMatch
	if a.r.Match(r, &m) {
		Log.Debug("matched route is: ", m.Route.GetName())
		Log.Debug("matched route handler is: ", m.Route.GetName())
	}

	next(w, r)
}