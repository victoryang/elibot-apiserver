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
	cmdSpeed = "speed"
	cmdSetMainfile = "set_mainfile"
	cmdCycleMode = "cycleMode"
)

func execCmd(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	args := vars["args"]

	Log.Debug("run ", args)
	cmd := ConcatCommand(cmdRun, args)
	SendToMCServerWithTimeout(w, r, cmd, TagExecute)
}

func execPause(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	args := vars["args"]

	Log.Debug("pause ", args)
	cmd := ConcatCommand(cmdPause, args)
	SendToMCServerWithTimeout(w, r, cmd, TagExecute)
}

func setRobotMode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mode := vars["mode"]

	Log.Debug("set mode ", mode)
	cmd := ConcatCommand(cmdMode, mode)
	SendToMCServerWithTimeout(w, r, cmd, TagExecute)
}

func clearAlarm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	args := vars["args"]

	Log.Debug("clear alarm ", args)
	cmd := ConcatCommand(cmdClearAlarm, args)
	SendToMCServerWithTimeout(w, r, cmd, TagExecute)
}

func setSpeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := vars["data"]

	Log.Debug("set speed ", data)
	cmd := ConcatCommand(cmdSpeed, data)
	SendToMCServerWithTimeout(w, r, cmd, TagExecute)
}

func setMainfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	Log.Debug("set main file ", filename)
	cmd := ConcatCommand(cmdSetMainfile, filename)
	SendToMCServerWithTimeout(w, r, cmd, TagExecute)
}

func setCycleMode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cyclemode := vars["cyclemode"]

	Log.Debug("set cycle mode ", cyclemode)
	cmd := ConcatCommand(cmdCycleMode, cyclemode)
	SendToMCServerWithTimeout(w, r, cmd, TagExecute)
}
