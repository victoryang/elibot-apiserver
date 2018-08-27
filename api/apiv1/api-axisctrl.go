package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	TagAxisCtrl = "apiv1:robot:axisctrl"

	cmdServo = "servo"
	cmdDragTeach = "drag_teach"
)

func setServoStatusv1 (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	status := vars["status"]

	Log.Debug("servo ", status)
	cmd := ConcatCommand(cmdServo, status)
	SendToMCServerWithTimeout(w, cmd, TagAxisCtrl)
}

func setDragteachStatus (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	status := vars["status"]

	Log.Debug("drag_teach ", status)
	cmd := ConcatCommand(cmdDragTeach, status)
	SendToMCServerWithTimeout(w, cmd, TagAxisCtrl)
}

func setServoStatus (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	status := vars["status"]

	Log.Debug("servo ", status)
	SendToMCServerWithJsonRpc(w, cmdServo, status)
}