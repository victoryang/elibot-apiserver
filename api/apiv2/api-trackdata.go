package apiv2

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	TagRepository = "apiv2:robot:trackdata"

	cmdSetTrackCraftNum = "set_trackcraftnum"
	cmdSetTrackdataNote = "cmd_set_trackdatanote"
	cmdSetTrackdata = "set_trackdata"
	cmdSaveTrackdata = "save_trackdata"
	cmdStartCalibration = "start_Calibration"
	cmdRecordPointA = "Recoder_PointA"
	cmdRecordPointB = "Recoder_PointB"
	cmdRecordPointC = "Recoder_PointC"
	cmdRecordReferencePoint = "Recoder_ReferencePoint"
	cmdGotoReferPos = "goto_referPos"
	cmdGetCalibrationRsult = "get_calibration_result"
)

func setTrackCraftNum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num := vars["num"]

	Log.Debug("set_trackcraftnum ", num)
	cmd := ConcatCommand(cmdSetTrackCraftNum, num)
	SendToMCServerWithTimeout(w, r, cmd, TagRepository)
}

func setTrackdataNote(w http.ResponseWriter, r *http.Request) {
	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}
	notes := strings.Join(d.Note, " ")
	Log.Debug("cmd_set_trackdatanote ", notes)
	cmd := ConcatCommand(cmdSetTrackdataNote, notes)
	SendToMCServerWithTimeout(w, r, cmd, TagUserFrame)
}

func setTrackdata(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	property := vars["property"]
	context := vars["context"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	Log.Debug("set_trackdata ", property, context, d.Value)
	cmd := ConcatCommand(cmdSetUserNum, , property, context, d.Value)
	SendToMCServerWithTimeout(w, r, cmd, TagUserFrame)
}

func saveTrackdata(w http.ResponseWriter, r *http.Request) {
	Log.Debug("save_trackdata ")
	SendToMCServerWithTimeout(w, r, cmdSaveTrackdata, TagUserFrame)
}

func startCalibration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trigger := vars["trigger"]
	Log.Debug("start_Calibration ", trigger)
	cmd := ConcatCommand(cmdStartCalibration, , trigger)
	SendToMCServerWithTimeout(w, r, cmd, TagUserFrame)
}

func recordPointA(w http.ResponseWriter, r *http.Request) {
	Log.Debug("recordPointA")
	SendToMCServerWithTimeout(w, r, cmdRecordPointA, TagUserFrame)
}

func recordPointB(w http.ResponseWriter, r *http.Request) {
	Log.Debug("recordPointB")
	SendToMCServerWithTimeout(w, r, cmdRecordPointB, TagUserFrame)
}

func recordPointC(w http.ResponseWriter, r *http.Request) {
	Log.Debug("recordPointC")
	SendToMCServerWithTimeout(w, r, cmdRecordPointC, TagUserFrame)
}

func recordReferencePoint(w http.ResponseWriter, r *http.Request) {
	Log.Debug("Recoder_ReferencePoint")
	SendToMCServerWithTimeout(w, r, cmdRecordReferencePoint, TagUserFrame)
}

func gotoReferPos(w http.ResponseWriter, r *http.Request) {
	Log.Debug("goto_referPos")
	SendToMCServerWithTimeout(w, r, cmdGotoReferPos, TagUserFrame)
}

func getCalibrationRsult(w http.ResponseWriter, r *http.Request) {
	Log.Debug("get_calibration_result")
	SendToMCServerWithTimeout(w, r, cmdGetCalibrationRsult, TagUserFrame)
}