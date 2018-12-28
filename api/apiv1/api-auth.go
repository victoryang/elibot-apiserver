package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	"elibot-apiserver/auth"
	Log "elibot-apiserver/log"
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

	var authority int
	if !auth.GetUserManager().GetUserAuthority(username, &authority) {
		Log.Debug("Can't get authority for this user")
		WriteUnauthorizedResponse(w)
		return
	}

	ip := getSourceIP(r)
	if isRemoteReq(ip) {
		if remoteModeOff() || !isRemoteAuthority(authority) {
			Log.Debug("remote check fails")
			WriteForbiddenResponse(w)
			return
		}
	}

	token := auth.SetSession(username, authority, ip)

	WriteSuccessResponse(w, token)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	auth.ClearSession(getTokenFromHeader(r))
	WriteSuccessResponse(w, "logout successfully\n")
}