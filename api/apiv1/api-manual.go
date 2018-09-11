package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	cmdCoord = "coord"
	cmdManual = "manual"
	cmdRunForward = "runForward"
	cmdRunToZero = "runToZero"
	cmdStop = "stop"
)

func setCoordinateMode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mode := vars["mode"]

	Log.Debug("set coord mode ", mode)
	SendToMCServerWithJsonRpc(w, cmdCoord, ConcatParams(mode))
}

func doManual(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	axis := vars["axis"]

	d := &RequestDataForCommandArgs{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteBadRequestResponse(w, ERRINVALIDBODY)
		return
	} 

	Log.Debug("set mannual ", axis, d.Args[:])
	var params []string
	params = append(params, axis)
	for _,v := range d.Args {
		params = append(params, v)
	}
	SendToMCServerWithJsonRpc(w, cmdManual, params)
}

func doRunForward(w http.ResponseWriter, r *http.Request) {
	d := &RequestDataForCommandArgs{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteBadRequestResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("run forward ", d.Args[:])
	SendToMCServerWithJsonRpc(w, cmdRunForward, ConcatParams(d.Args...))
}

func doRunToZero(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]

	Log.Debug("runToZero ", status)
	SendToMCServerWithJsonRpc(w, cmdRunToZero, ConcatParams(status))
}

func doRobotStop(w http.ResponseWriter, r *http.Request) {
	Log.Debug("robot stop")
	SendToMCServerWithJsonRpc(w, cmdStop, nil)
}