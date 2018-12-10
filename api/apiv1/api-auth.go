package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	"elibot-apiserver/auth"
	//Log "elibot-apiserver/log"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["username"] == "" || vars["pwd"] == "" {
		WriteBadRequestResponse(w, ERR_REQ_INVALID_PARAMETER)
		return
	}

	username := vars["username"]
	if !auth.GetUserManager().VerifyPassword(username, vars["pwd"]) {
		WriteUnauthorizedResponse(w)
		return
	}

	ip := getSourceIp(r)
	if auth.CheckSession(ip) {
		WriteForbiddenResponse(w)
		return
	}

	token := auth.SetSession(username, ip)

	WriteSuccessResponse(w, token)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	auth.ClearSession(getSourceIp(r))
	WriteSuccessResponse(w, "logout successfully\n")
}