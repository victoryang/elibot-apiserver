package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	TagVisionData = "apiv1:robot:visiondata"

	cmdSetVisionCraftNum = "set_visioncraftnum"
	cmdSetVisionDataNote = "set_visiondatanote"
	cmdSetVisionData = "set_visiondata"
	cmdSaveVisionData = "save_visiondata"
	cmdTriggerCamera = "Camera_trigger"
	cmdRecordVisionEncoVal = "recode_VisionEncoVal"
	cmdSetVisualCalibration = "Visual_calibration"
	cmdRecordRobEncoVal = "recode_RobEncoVal"
	cmdRecordRobotPoint = "recode_RobotPoint"
	cmdRecordVisualPoint = "recode_VisualPoint"
	cmdTakePhoto = "take_photo"
	cmdDataRefresh = "Data_refresh"
	cmdGoToPoint = "goto_ThisPoint"
)

type Note struct {
	Value  		[]string		`json:"value,omitempty"`
}

func setVisionCraftNum(w http.ResponseWriter, r *http.Request) {
	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("set_visioncraftnum ", d.Value)
	cmd := ConcatCommand(cmdSetVisionCraftNum, d.Value)
	SendToMCServerWithTimeout(w, cmd, TagVisionData)
}

func setVisionDataNote(w http.ResponseWriter, r *http.Request) {
	n := &Note{}
	if err := ParseBodyToObject(r, n); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("set_visiondatanote ", n.Value)
	cmd := ConcatCommand(cmdSetVisionDataNote, n.Value...)
	SendToMCServerWithTimeout(w, cmd, TagVisionData)
}

func setVisionData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	property := vars["property"]
	context := vars["context"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("set_visiondata ", property, context, d.Value)
	cmd := ConcatCommand(cmdSetVisionData, property, context, d.Value)
	SendToMCServerWithTimeout(w, cmd, TagVisionData)
}

func saveVisionData(w http.ResponseWriter, r *http.Request) {
	Log.Debug("save_visiondata")
	SendToMCServerWithTimeout(w, cmdSaveVisionData, TagVisionData)
}

func triggerCamera(w http.ResponseWriter, r *http.Request) {
	Log.Debug("Camera_trigger")
	SendToMCServerWithTimeout(w, cmdTriggerCamera, TagVisionData)
}

func recordVisionEncoVal(w http.ResponseWriter, r *http.Request) {
	Log.Debug("recode_VisionEncoVal")
	SendToMCServerWithTimeout(w, cmdRecordVisionEncoVal, TagVisionData)
}

func setVisualCalibration(w http.ResponseWriter, r *http.Request) {
	Log.Debug("Visual_calibration")
	SendToMCServerWithTimeout(w, cmdSetVisualCalibration, TagVisionData)
}

func recordRobEncoVal(w http.ResponseWriter, r *http.Request) {
	Log.Debug("recode_RobEncoVal")
	SendToMCServerWithTimeout(w, cmdRecordRobEncoVal, TagVisionData)
}

func recordRobotPoint(w http.ResponseWriter, r *http.Request) {
	Log.Debug("recode_RobotPoint")
	SendToMCServerWithTimeout(w, cmdRecordRobotPoint, TagVisionData)
}

func recordVisualPoint(w http.ResponseWriter, r *http.Request) {
	Log.Debug("recode_VisualPoint")
	SendToMCServerWithTimeout(w, cmdRecordVisualPoint, TagVisionData)
}

func takePhoto(w http.ResponseWriter, r *http.Request) {
	Log.Debug("take_photo")
	SendToMCServerWithTimeout(w, cmdTakePhoto, TagVisionData)
}

func dataReresh(w http.ResponseWriter, r *http.Request) {
	Log.Debug("Data_refresh")
	SendToMCServerWithTimeout(w, cmdDataRefresh, TagVisionData)
}

func goToPoint(w http.ResponseWriter, r *http.Request) {
	Log.Debug("goto_ThisPoint")
	SendToMCServerWithTimeout(w, cmdGoToPoint, TagVisionData)
}