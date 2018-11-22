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
	cmdSetCurLine = "set_curLine"
	cmdGobackOrigin = "getOrigin"
)

func doRunCmd(w http.ResponseWriter, r *http.Request) {
	d := &RequestDataForCommandArgs{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteBadRequestResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("run ", d.Args[:])
	SendToMCServer(w, cmdRun, ConcatParams(d.Args...))
}

func doPause(w http.ResponseWriter, r *http.Request) {
	Log.Debug("pause")
	SendToMCServer(w, cmdPause, nil)
}

func setRobotMode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mode := vars["mode"]

	Log.Debug("set mode ", mode)
	SendToMCServer(w, cmdMode, ConcatParams(mode))
}

func doClearAlarm(w http.ResponseWriter, r *http.Request) {
	d := &RequestDataForCommandArgs{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteBadRequestResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("clear alarm ", d.Args[:])
	SendToMCServer(w, cmdClearAlarm, ConcatParams(d.Args...))
}

func doProgReset(w http.ResponseWriter, r *http.Request) {
	Log.Debug("progReset")
	SendToMCServer(w, cmdResetProg, nil)
}

func setSpeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := vars["data"]

	Log.Debug("set speed ", data)
	SendToMCServer(w, cmdSpeed, ConcatParams(data))
}

func setMainfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	progname := vars["progname"]

	Log.Debug("set main file ", progname)
	SendToMCServer(w, cmdSetMainfile, ConcatParams(progname))
}

func setCycleMode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cyclemode := vars["cyclemode"]

	Log.Debug("set cycle mode ", cyclemode)
	SendToMCServer(w, cmdCycleMode, ConcatParams(cyclemode))
}

func doLoadFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	Log.Debug("load file ", filename)
	SendToMCServer(w, cmdLoadFile, ConcatParams(filename))
}

func setCurLine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lineno := vars["lineno"]

	Log.Debug("set cur line ", lineno)
	SendToMCServer(w, cmdSetCurLine, ConcatParams(lineno))
}

func doGobackOrigin(w http.ResponseWriter, r *http.Request) {
	Log.Debug("gobackorigin")
	SendToMCServer(w, cmdGobackOrigin, nil)
}