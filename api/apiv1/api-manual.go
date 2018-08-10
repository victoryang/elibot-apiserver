package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	TagManualInterpolation = "apiv1:robot:manualinterpolation"

	cmdCoord = "coord"
	cmdManual = "manual "
	cmdRunForward = "runForward"
	cmdRunToZero = "runToZero"
	cmdStop = "stop"
)

func setCoordinateMode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mode := vars["mode"]

	Log.Debug("set coord mode ", mode)
	cmd := ConcatCommand(cmdCoord, mode)
	SendToMCServerWithTimeout(w, r, cmd, TagManualInterpolation)
}

func doManual(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	axis := vars["axis"]

	d := &RequestDataForCommandArgs{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	} 

	Log.Debug("set mannual ", axis, d.Args[:])
	cmdtmp := cmdManual + axis
	cmd := ConcatCommand(cmdtmp, d.Args...)
	SendToMCServerWithTimeout(w, r, cmd, TagManualInterpolation)
}

func doRunForward(w http.ResponseWriter, r *http.Request) {
	d := &RequestDataForCommandArgs{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("run forward ", d.Args[:])
	cmd := ConcatCommand(cmdRunForward, d.Args...)
	SendToMCServerWithTimeout(w, r, cmd, TagManualInterpolation)
}

func doRunToZero(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]

	Log.Debug("runToZero ", status)
	cmd := ConcatCommand(cmdRunToZero, status)
	SendToMCServerWithTimeout(w, r, cmd, TagManualInterpolation)
}

func doRobotStop(w http.ResponseWriter, r *http.Request) {
	Log.Debug("robot stop")
	SendToMCServerWithTimeout(w, r, cmdStop, TagManualInterpolation)
}