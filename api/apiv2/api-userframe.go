package apiv2

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	TagUserFrame = "apiv2:robot:userframe"

	cmd_SetUserFrame = "set_userframe"
	cmdSetUserPos = "set_userpos"
	cmdGoToUserPos = "goto_userpos"
	cmdSetUserNum= "set_usernum"
	cmdSetUserNote = "set_usernote"
)

func set_UserFrame(w http.ResponseWriter, r *http.Request) {
	Log.Debug("set_userframe ")
	SendToMCServerWithTimeout(w, r, cmd_SetUserFrame, TagUserFrame)
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

func setUserNote(w http.ResponseWriter, r *http.Request) {
	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}
	notes := strings.Join(d.Note, " ")
	Log.Debug("set_usernote ", notes)
	cmd := ConcatCommand(cmdSetUserNote, notes)
	SendToMCServerWithTimeout(w, r, cmd, TagUserFrame)
}