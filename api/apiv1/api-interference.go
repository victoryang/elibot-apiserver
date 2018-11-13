package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	cmdSetInterferData = "set_interf_data"
)

func setInterferData (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	num := vars["num"]

	Log.Debug("set_interf_data ", num)
	SendToMCServerWithJsonRpc(w, cmdSetInterferData, ConcatParams(num))
}