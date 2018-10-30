package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "api-server/log"
)

const (
	cmdSetUserPos = "set_userpos"
	cmdGotoUserPos = "goto_userpos"
	cmdCalUserFrame = "set_userframe"
)

func setUserPos (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	pos_no := vars["pos_no"]

	Log.Debug("set_userpos ", pos_no)
	SendToMCServerWithJsonRpc(w, cmdSetUserPos, ConcatParams(pos_no))
}

func doGotoUserPos (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	pos_no := vars["pos_no"]

	Log.Debug("goto_userpos ", pos_no)
	SendToMCServerWithJsonRpc(w, cmdGotoUserPos, ConcatParams(pos_no))
}

func doCalUserFrame (w http.ResponseWriter, r *http.Request){
	Log.Debug("set_userframe")
	SendToMCServerWithJsonRpc(w, cmdCalUserFrame, nil)
}