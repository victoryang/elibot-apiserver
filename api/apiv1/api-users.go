package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	"elibot-apiserver/auth"
	"elibot-apiserver/db"
	//Log "elibot-apiserver/log"
)

const (
	ADMINISTRATOR = "Admin"
)

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
	if u.Name == ADMINISTRATOR || u.Authority == AuthorityAdmin {
		WriteForbiddenResponse(w)
		return
	}

	if !auth.GetUserManager().AddUser(u, vars["pwd"]) {
		WriteInternalServerErrorResponse(w, ERRFAILTOOPERATEUSER)
		return
	}

	WriteSuccessResponse(w, "Add user successfully\n")
}

func removeUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["username"] == ADMINISTRATOR {
		WriteForbiddenResponse(w)
		return
	}

	if !auth.GetUserManager().RemoveUser(vars["username"]) {
		WriteInternalServerErrorResponse(w, ERRFAILTOOPERATEUSER)
		return
	}

	WriteSuccessResponse(w, "Remove user successfully\n")
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

	if !auth.GetUserManager().VerifyPassword(vars["username"], vars["pwd"]) {
		WriteUnauthorizedResponse(w)
		return
	}

	u.Name = vars["username"]
	if !auth.GetUserManager().ModifyUser(u) {
		WriteInternalServerErrorResponse(w, ERRFAILTOOPERATEUSER)
		return
	}

	WriteSuccessResponse(w, "Modify user successfully\n")
}

func changePassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["username"] == "" || vars["dpwd"] == "" {
		WriteBadRequestResponse(w, ERR_REQ_INVALID_PARAMETER)
		return
	}

	token := getTokenFromHeader(r)
	if !auth.CheckSession(token) {
		WriteUnauthorizedResponse(w)
		return
	}

	cu := auth.GetLoginedUserInfo(token)
	if cu.Username != ADMINISTRATOR && cu.Username != vars["username"] {
		WriteForbiddenResponse(w)
		return
	}

	if cu.Username == vars["username"] {
		if !auth.GetUserManager().VerifyPassword(vars["username"], vars["spwd"]) {
			WriteUnauthorizedResponse(w)
			return
		}
	}

	if !auth.GetUserManager().ChangePassword(vars["username"], vars["dpwd"]) {
		WriteInternalServerErrorResponse(w, ERRFAILTOOPERATEPWD)
		return
	}

	WriteSuccessResponse(w, "Change password successfully\n")
}

func verifyPwdOnceLogined(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ip := getSourceIP(r)
	if !auth.CheckSession(ip) {
		WriteForbiddenResponse(w)
		return
	}

	if vars["username"] == "" || vars["pwd"] == "" {
		WriteBadRequestResponse(w, ERR_REQ_INVALID_PARAMETER)
		return
	}

	username := vars["username"]
	if !auth.GetUserManager().VerifyPassword(username, vars["pwd"]) {
		WriteUnauthorizedResponse(w)
		return
	}

	WriteSuccessResponse(w, true)
}