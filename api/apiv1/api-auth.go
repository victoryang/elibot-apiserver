package apiv1

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"elibot-apiserver/auth"
	"elibot-apiserver/db"
	//Log "elibot-apiserver/log"
)

func getSourceIp(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		WriteUnauthorizedResponse(w)
		return
	}

	if res := db.VerifySecretkey(username, password); !res {
		WriteUnauthorizedResponse(w)
		return
	}

	token := auth.SetSession(username, getSourceIp(r))

	cookie := &http.Cookie {
		Name: 	"session",
		Value: 	token,
	}
	http.SetCookie(w, cookie)
	WriteSuccessResponse(w, "login successfully")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	auth.ClearSession(getSourceIp(r))
	WriteSuccessResponse(w, "logout successfully")
}

func getUserList(w http.ResponseWriter, r *http.Request) {
	WriteSuccessResponse(w, auth.GetUserManager().GetUsersList())
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var u db.User
	auth.GetUserManager().GetUser(vars["username"], &u)
	WriteSuccessResponse(w, u)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	u := db.User{}
	if err := ParseBodyToObject(r, &u); err!=nil {
		WriteBadRequestResponse(w, ERRINVALIDBODY)
		return
	}

	u.Name = vars["username"]
	success := auth.GetUserManager().AddUser(u, vars["password"])
	WriteSuccessResponse(w, success)
}

func removeUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	success := auth.GetUserManager().RemoveUser(vars["username"])
	WriteSuccessResponse(w, success)
}

func modifyUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	u := db.User{}
	if err := ParseBodyToObject(r, &u); err!=nil {
		WriteBadRequestResponse(w, ERRINVALIDBODY)
		return
	}

	u.Name = vars["username"]
	success := auth.GetUserManager().ModifyUser(u)
	WriteSuccessResponse(w, success)
}

func changePassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	success := auth.GetUserManager().ChangePassword(vars["username"], vars["password"])
	WriteSuccessResponse(w, success)
}