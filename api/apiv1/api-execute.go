package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	TagExecute = "apiv1:robot:execute"

	cmdRun = "run"
	cmdPause = "pause"
	cmdMode = "mode"
	cmdClearAlarm = "ClearAlarm"
	cmdResetProg = "progReset"
	cmdSpeed = "speed"
	cmdSetMainfile = "set_mainfile"
	cmdCycleMode = "cycleMode"
)

func doRunCmd(w http.ResponseWriter, r *http.Request) {
	d := &RequestDataForCommandArgs{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("run ", d.Args[:])
	cmd := ConcatCommand(cmdRun, d.Args...)
	SendToMCServerWithTimeout(w, cmd, TagExecute)
}

func doPause(w http.ResponseWriter, r *http.Request) {
	SendToMCServerWithTimeout(w, cmdPause, TagExecute)
}

func setRobotMode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mode := vars["mode"]

	Log.Debug("set mode ", mode)
	cmd := ConcatCommand(cmdMode, mode)
	SendToMCServerWithTimeout(w, cmd, TagExecute)
}

func doClearAlarm(w http.ResponseWriter, r *http.Request) {
	d := &RequestDataForCommandArgs{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("clear alarm ", d.Args[:])
	cmd := ConcatCommand(cmdClearAlarm, d.Args...)
	SendToMCServerWithTimeout(w, cmd, TagExecute)
}

func doProgReset(w http.ResponseWriter, r *http.Request) {
	Log.Debug("progReset")
	SendToMCServerWithTimeout(w, cmdResetProg, TagExecute)
}

func setSpeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := vars["data"]

	Log.Debug("set speed ", data)
	cmd := ConcatCommand(cmdSpeed, data)
	SendToMCServerWithTimeout(w, cmd, TagExecute)
}

func setMainfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	Log.Debug("set main file ", filename)
	cmd := ConcatCommand(cmdSetMainfile, filename)
	SendToMCServerWithTimeout(w, cmd, TagExecute)
}

func setCycleMode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cyclemode := vars["cyclemode"]

	Log.Debug("set cycle mode ", cyclemode)
	cmd := ConcatCommand(cmdCycleMode, cyclemode)
	SendToMCServerWithTimeout(w, cmd, TagExecute)
}
