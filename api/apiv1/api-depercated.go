package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	TagAxisCtrl = "apiv1:robot:axisctrl"
	TagExecute = "apiv1:robot:execute"
	TagManualInterpolation = "apiv1:robot:manual"
	TagRepository = "apiv1:robot:repository"
)

func setServoStatus_deprecated (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	status := vars["status"]

	Log.Debug("servo ", status)
	cmd := ConcatCommand(cmdServo, status)
	SendToMCServerWithTimeout(w, cmd, TagAxisCtrl)
}

func setDragteachStatus_deprecated (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	status := vars["status"]

	Log.Debug("drag_teach ", status)
	cmd := ConcatCommand(cmdDragTeach, status)
	SendToMCServerWithTimeout(w, cmd, TagAxisCtrl)
}

func doRunCmd_deprecated (w http.ResponseWriter, r *http.Request) {
	d := &RequestDataForCommandArgs{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("run ", d.Args[:])
	cmd := ConcatCommand(cmdRun, d.Args...)
	SendToMCServerWithTimeout(w, cmd, TagExecute)
}

func doPause_deprecated (w http.ResponseWriter, r *http.Request) {
	SendToMCServerWithTimeout(w, cmdPause, TagExecute)
}

func setRobotMode_deprecated (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mode := vars["mode"]

	Log.Debug("set mode ", mode)
	cmd := ConcatCommand(cmdMode, mode)
	SendToMCServerWithTimeout(w, cmd, TagExecute)
}

func doClearAlarm_deprecated (w http.ResponseWriter, r *http.Request) {
	d := &RequestDataForCommandArgs{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("clear alarm ", d.Args[:])
	cmd := ConcatCommand(cmdClearAlarm, d.Args...)
	SendToMCServerWithTimeout(w, cmd, TagExecute)
}

func doProgReset_deprecated (w http.ResponseWriter, r *http.Request) {
	Log.Debug("progReset")
	SendToMCServerWithTimeout(w, cmdResetProg, TagExecute)
}

func setSpeed_deprecated (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := vars["data"]

	Log.Debug("set speed ", data)
	cmd := ConcatCommand(cmdSpeed, data)
	SendToMCServerWithTimeout(w, cmd, TagExecute)
}

func setMainfile_deprecated (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	Log.Debug("set main file ", filename)
	cmd := ConcatCommand(cmdSetMainfile, filename)
	SendToMCServerWithTimeout(w, cmd, TagExecute)
}

func setCycleMode_deprecated (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cyclemode := vars["cyclemode"]

	Log.Debug("set cycle mode ", cyclemode)
	cmd := ConcatCommand(cmdCycleMode, cyclemode)
	SendToMCServerWithTimeout(w, cmd, TagExecute)
}

func setCoordinateMode_deprecated (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mode := vars["mode"]

	Log.Debug("set coord mode ", mode)
	cmd := ConcatCommand(cmdCoord, mode)
	SendToMCServerWithTimeout(w, cmd, TagManualInterpolation)
}

func doManual_deprecated (w http.ResponseWriter, r *http.Request) {
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
	SendToMCServerWithTimeout(w, cmd, TagManualInterpolation)
}

func doRunForward_deprecated (w http.ResponseWriter, r *http.Request) {
	d := &RequestDataForCommandArgs{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("run forward ", d.Args[:])
	cmd := ConcatCommand(cmdRunForward, d.Args...)
	SendToMCServerWithTimeout(w, cmd, TagManualInterpolation)
}

func doRunToZero_deprecated (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]

	Log.Debug("runToZero ", status)
	cmd := ConcatCommand(cmdRunToZero, status)
	SendToMCServerWithTimeout(w, cmd, TagManualInterpolation)
}

func doRobotStop_deprecated (w http.ResponseWriter, r *http.Request) {
	Log.Debug("robot stop")
	SendToMCServerWithTimeout(w, cmdStop, TagManualInterpolation)
}

func setArcParam_deprecated (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	file_no := vars["file_no"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("setArcParam ", md_id, d.Value, file_no, strconv.Itoa(d.Index))
	cmd := ConcatCommand(cmdSetArcParam, md_id, d.Value, file_no, strconv.Itoa(d.Index))
	SendToMCServerWithTimeout(w, cmd, TagRepository)
}

func setInterference_deprecated (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	no := vars["no"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("setInterference ", md_id, d.Value, no, strconv.Itoa(d.Index))
	cmd := ConcatCommand(cmdSetInterference, md_id, d.Value, no, strconv.Itoa(d.Index))
	SendToMCServerWithTimeout(w, cmd, TagRepository)
}

func setParam_deprecated (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("setParam ", md_id, d.Value, strconv.Itoa(d.Index))
	cmd := ConcatCommand(cmdSetParam, md_id, d.Value, strconv.Itoa(d.Index))
	SendToMCServerWithTimeout(w, cmd, TagRepository)
}

func setToolFrame_deprecated (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	tool_no := vars["tool_no"]
	pos_no := vars["pos_no"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("setToolFrame ", md_id, d.Value, tool_no, pos_no, strconv.Itoa(d.Index))
	cmd := ConcatCommand(cmdSetToolFrame, md_id, d.Value, tool_no, pos_no, strconv.Itoa(d.Index))
	SendToMCServerWithTimeout(w, cmd, TagRepository)
}

func setUserFrame_deprecated (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	userno := vars["userno"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("setUserFrame ", md_id, d.Value, userno)
	cmd := ConcatCommand(cmdSetUserFrame, md_id, d.Value, userno)
	SendToMCServerWithTimeout(w, cmd, TagRepository)
}

func setZeroPoint_deprecated (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("setZeroPoint ", md_id, d.Value, strconv.Itoa(d.Index))
	cmd := ConcatCommand(cmdSetZeroPoint, md_id, d.Value, strconv.Itoa(d.Index))
	SendToMCServerWithTimeout(w, cmd, TagRepository)
}