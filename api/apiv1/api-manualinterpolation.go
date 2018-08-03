package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	TagManualInterpolation = "apiv1:robot:manualinterpolation"

	cmdCoord = "coord"
	cmdStop = "stop"
	cmdManual = "manual"
	cmdRunForward = "runForward"
	cmdRunToZero = "runToZero"
)

func setCoordinateMode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mode := vars["mode"]

	Log.Debug("set coord mode ", mode)
	cmd := ConcatCommand(cmdCoord, mode)
	SendToMCServerWithTimeout(w, r, cmd, TagManualInterpolation)
}

func robotStop(w http.ResponseWriter, r *http.Request) {
	Log.Debug("robot stop")
	SendToMCServerWithTimeout(w, r, cmdStop, TagManualInterpolation)
}

func setManual(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	arg := vars["arg"]

	Log.Debug("set mannual ", key, arg)
	cmd := ConcatCommand(cmdManual, key, arg)
	SendToMCServerWithTimeout(w, r, cmd, TagManualInterpolation)
}

func runForward(w http.ResponseWriter, r *http.Request) {
	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("run forward ", d.Values[:])
	cmd := ConcatCommand(cmdRunForward, d.Values...)
	SendToMCServerWithTimeout(w, r, cmd, TagManualInterpolation)
}

func runToZero(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	zero := vars["zero"]

	Log.Debug("run to zero ", zero)
	cmd := ConcatCommand(cmdRunToZero, zero)
	SendToMCServerWithTimeout(w, r, cmd, TagManualInterpolation)
}