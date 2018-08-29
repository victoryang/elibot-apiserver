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

	d := &RequestDataForCommandArgs{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	var args []string
	if d.Args != nil {
		args = append(ConcatParams(status), d.Args...)
	} else {
		args = ConcatParams(status)
	}

	Log.Debug("servo ", status)
	SendToMCServerWithJsonRpc(w, cmdServo, args)
}

func setDragteachStatus (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	status := vars["status"]

	Log.Debug("drag_teach ", status)
	SendToMCServerWithJsonRpc(w, cmdDragTeach, ConcatParams(status))
}