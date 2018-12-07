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
	vars := mux.Vars(r)

	if vars["username"] == "" || vars["pwd"] == "" {
		WriteBadRequestResponse(w, ERR_REQ_INVALID_PARAMETER)
		return
	}

	username := vars["username"]
	if res := auth.GetUserManager().VerifyPassword(username, vars["pwd"]); !res {
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
	WriteSuccessResponse(w, "logout successfully")
}

func getUserList(w http.ResponseWriter, r *http.Request) {
	start, end, err := validateRange(r)
	if err != nil {
		WriteBadRequestResponse(w, ERRINCORRECTRANGE)
		return
	}

	users := auth.GetUserManager().GetUsersList()
	if start > len(users) {
		WriteBadRequestResponse(w, ERRINCORRECTRANGE)
		return
	} else if end > len(users) {
		end = len(users)
	}

	type UserList struct{
		Users 			[]db.User 	`json:"users"`
		TotalPageSize 	int 		`json:"totalPageSize"`
	}

	res := users[start:end]

	WriteSuccessResponse(w, UserList{Users: res, TotalPageSize: len(users)})
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var u db.User
	if !auth.GetUserManager().GetUser(vars["username"], &u) {
		WriteInternalServerErrorResponse(w, ERRFAILTOLISTUSER)
		return
	}

	WriteSuccessResponse(w, u)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	u := db.User{}
	if err := ParseBodyToObject(r, &u); err!=nil {
		WriteBadRequestResponse(w, ERRINVALIDBODY)
		return
	}

	if vars["username"] == "" || vars["pwd"] == "" {
		WriteBadRequestResponse(w, ERR_REQ_INVALID_PARAMETER)
		return
	}

	u.Name = vars["username"]

	if !auth.GetUserManager().AddUser(u, vars["pwd"]) {
		WriteInternalServerErrorResponse(w, ERRFAILTOADDUSER)
		return
	}

	WriteSuccessResponse(w, "Add user successfully")
}

func removeUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !auth.GetUserManager().RemoveUser(vars["username"]) {
		WriteInternalServerErrorResponse(w, ERRFAILTOREMOVEUSER)
		return
	}

	WriteSuccessResponse(w, "Remove user successfully")
}

func modifyUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	u := db.User{}
	if err := ParseBodyToObject(r, &u); err!=nil {
		WriteBadRequestResponse(w, ERRINVALIDBODY)
		return
	}

	if vars["username"] == "" || vars["pwd"] == "" {
		WriteBadRequestResponse(w, ERR_REQ_INVALID_PARAMETER)
		return
	}

	if res := auth.GetUserManager().VerifyPassword(vars["username"], vars["pwd"]); !res {
		WriteUnauthorizedResponse(w)
		return
	}

	u.Name = vars["username"]
	if !auth.GetUserManager().ModifyUser(u) {
		WriteInternalServerErrorResponse(w, ERRFAILTOMODIFYUSER)
		return
	}

	WriteSuccessResponse(w, "Modify user successfully")
}

func changePassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["username"] == "" || vars["spwd"] == "" || vars["dpwd"] == "" {
		WriteBadRequestResponse(w, ERR_REQ_INVALID_PARAMETER)
		return
	}

	if !auth.GetUserManager().ChangePassword(vars["username"], vars["spwd"], vars["dpwd"]) {
		WriteInternalServerErrorResponse(w, ERRFAILTOCHANGEPWD)
		return
	}

	WriteSuccessResponse(w, "Change password successfully")
}