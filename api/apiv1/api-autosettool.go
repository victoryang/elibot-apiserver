package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	cmdGotoToolPos = "goto_toolpos"
	cmdAutoSetToolFrame = "auto_set_toolframe"
)

func doGotoToolPos (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	num := vars["num"]

	Log.Debug("goto_toolpos ", num)
	SendToMCServerWithJsonRpc(w, cmdGotoToolPos, ConcatParams(num))
}

func doAutoSetToolFrame (w http.ResponseWriter, r *http.Request){
	Log.Debug("auto_set_toolframe")
	SendToMCServerWithJsonRpc(w, cmdAutoSetToolFrame, nil)
}