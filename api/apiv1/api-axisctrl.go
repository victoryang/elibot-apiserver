package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	cmdServo = "servo"
	cmdDragTeach = "drag_teach"
)

func setServoStatus (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	status := vars["status"]

	Log.Debug("servo ", status)
	SendToMCServerWithJsonRpc(w, cmdServo, ConcatParams(status))
}

func setDragteachStatus (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	status := vars["status"]

	Log.Debug("drag_teach ", status)
	SendToMCServerWithJsonRpc(w, cmdDragTeach, ConcatParams(status))
}