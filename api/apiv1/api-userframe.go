package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	cmdSetUserPos = "set_userpos"
	cmdGotoUserPos = "goto_userpos"
	cmdCalUserFrame = "set_userframe"
)

func setUserPos (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	user_no := vars["user_no"]
	pos_no := vars["pos_no"]

	Log.Debug("set_userpos ", user_no, pos_no)
	SendToMCServer(w, cmdSetUserPos, ConcatParams(user_no, pos_no))
}

func doGotoUserPos (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	user_no := vars["user_no"]
	pos_no := vars["pos_no"]

	Log.Debug("goto_userpos ", user_no, pos_no)
	SendToMCServer(w, cmdGotoUserPos, ConcatParams(user_no, pos_no))
}

func doCalUserFrame (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	user_no := vars["user_no"]

	Log.Debug("set_userframe ", user_no)
	SendToMCServer(w, cmdCalUserFrame, ConcatParams(user_no))
}