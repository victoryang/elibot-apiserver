package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	cmdRun = "run"
	cmdPause = "pause"
	cmdMode = "mode"
	cmdClearAlarm = "ClearAlarm"
	cmdResetProg = "progReset"
	cmdSpeed = "speed"
	cmdSetMainfile = "set_mainfile"
	cmdCycleMode = "cycleMode"
	cmdLoadFile = "load"
)

func doRunCmd(w http.ResponseWriter, r *http.Request) {
	d := &RequestDataForCommandArgs{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteBadRequestResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("run ", d.Args[:])
	SendToMCServerWithJsonRpc(w, cmdRun, ConcatParams(d.Args...))
}

func doPause(w http.ResponseWriter, r *http.Request) {
	Log.Debug("pause")
	SendToMCServerWithJsonRpc(w, cmdPause, nil)
}

func setRobotMode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mode := vars["mode"]

	Log.Debug("set mode ", mode)
	SendToMCServerWithJsonRpc(w, cmdMode, ConcatParams(mode))
}

func doClearAlarm(w http.ResponseWriter, r *http.Request) {
	d := &RequestDataForCommandArgs{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteBadRequestResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("clear alarm ", d.Args[:])
	SendToMCServerWithJsonRpc(w, cmdClearAlarm, ConcatParams(d.Args...))
}

func doProgReset(w http.ResponseWriter, r *http.Request) {
	Log.Debug("progReset")
	SendToMCServerWithJsonRpc(w, cmdResetProg, nil)
}

func setSpeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := vars["data"]

	Log.Debug("set speed ", data)
	SendToMCServerWithJsonRpc(w, cmdSpeed, ConcatParams(data))
}

func setMainfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	progname := vars["progname"]

	Log.Debug("set main file ", progname)
	SendToMCServerWithJsonRpc(w, cmdSetMainfile, ConcatParams(progname))
}

func setCycleMode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cyclemode := vars["cyclemode"]

	Log.Debug("set cycle mode ", cyclemode)
	SendToMCServerWithJsonRpc(w, cmdCycleMode, ConcatParams(cyclemode))
}

func doLoadFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	Log.Debug("load file ", filename)
	SendToMCServerWithJsonRpc(w, cmdLoadFile, ConcatParams(filename))
}