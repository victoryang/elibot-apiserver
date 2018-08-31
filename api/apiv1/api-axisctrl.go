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
		WriteBadRequestResponse(w, ERRINVALIDBODY)
		return
	}

	var params []string
	params = append(params, status)
	for _,v := range d.Args {
		params = append(params, v)
	}

	Log.Debug("servo ", status)
	SendToMCServerWithJsonRpc(w, cmdServo, params)
}

func setDragteachStatus (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	status := vars["status"]

	Log.Debug("drag_teach ", status)
	SendToMCServerWithJsonRpc(w, cmdDragTeach, ConcatParams(status))
}