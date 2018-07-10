package apiv2

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	TagAxisCtrl = "apiv2:robot:robcalibtrate"

	cmdAutoCalcZero = "auto_calc_zero"
	cmdAutoRunZero = "auto_run_zero"
)

func setAutoCalcZero(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	Log.Debug("auto_calc_zero ", filename)
	cmd := ConcatCommand(cmdAutoCalcZero, filename)
	SendToMCServerWithTimeout(w, r, cmd, TagAxisCtrl)
}

func setAutoRunZero(w http.ResponseWriter, r *http.Request) {
	Log.Debug("auto_run_zero")
	SendToMCServerWithTimeout(w, r, cmdAutoRunZero, TagAxisCtrl)
}