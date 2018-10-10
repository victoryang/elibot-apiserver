package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "api-server/log"
)

const (
	cmdSetUserPos = "set_userpos"
)

func setUserPos (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	pos_no := vars["pos_no"]

	Log.Debug("set_userpos ", pos_no)
	SendToMCServerWithJsonRpc(w, cmdSetUserPos, ConcatParams(pos_no))
}