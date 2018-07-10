package apiv2

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	TagUserFrame = "apiv2:robot:userframe"

	cmdSetUserFrame = "set_userframe"
	cmdSetUserPos = "set_userpos"
	cmdGoToUserPos = "goto_userpos"
	cmdSetUserNum= "set_usernum"
)

func setUserFrame(w http.ResponseWriter, r *http.Request) {
	Log.Debug("set_userframe ")
	SendToMCServerWithTimeout(w, r, cmdSetUserFrame, TagUserFrame)
}

func setUserPos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pos_no := vars["pos_no"]

	Log.Debug("set_userpos ",pos_no)
	cmd := ConcatCommand(cmdSetUserPos, pos_no)
	SendToMCServerWithTimeout(w, r, cmd, TagUserFrame)
}

func goToUserPos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pos_no := vars["pos_no"]

	Log.Debug("goto_userpos ",pos_no)
	cmd := ConcatCommand(cmdSetUserPos, pos_no)
	SendToMCServerWithTimeout(w, r, cmd, TagUserFrame)
}

func setUserNum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num := vars["num"]

	Log.Debug("set_usernum ",num)
	cmd := ConcatCommand(cmdSetUserNum, num)
	SendToMCServerWithTimeout(w, r, cmd, TagUserFrame)
}

/*func setUserNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num := vars["num"]

	Log.Debug("set_usernum ",num)
	cmd := ConcatCommand(cmdSetUserNum, num)
	SendToMCServerWithTimeout(w, r, cmd, TagUserFrame)
}*/