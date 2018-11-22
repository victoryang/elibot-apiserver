package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	cmdGotoToolPos = "goto_toolpos"
	cmdAutoSetToolFrame = "auto_set_toolframe"
	cmdSetToolPos = "set_toolpos"
	cmdClearToolPos = "clear_toolpos"
)

func doGotoToolPos (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	tool_no := vars["tool_no"]
	num := vars["num"]

	Log.Debug("goto_toolpos ", tool_no, num)
	SendToMCServer(w, cmdGotoToolPos, ConcatParams(tool_no, num))
}

func doAutoSetToolFrame (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	tool_no := vars["tool_no"]

	Log.Debug("auto_set_toolframe ", tool_no)
	SendToMCServer(w, cmdAutoSetToolFrame, ConcatParams(tool_no))
}

func doSetToolPos (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	tool_no := vars["tool_no"]
	num := vars["num"]

	Log.Debug("set_toolpos ", tool_no, num)
	SendToMCServer(w, cmdSetToolPos, ConcatParams(tool_no, num))
}

func doClearToolPos (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	tool_no := vars["tool_no"]
	num := vars["num"]

	Log.Debug("set_toolpos ", tool_no, num)
	SendToMCServer(w, cmdClearToolPos, ConcatParams(tool_no, num))
}