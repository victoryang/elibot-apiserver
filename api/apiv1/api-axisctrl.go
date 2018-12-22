package apiv1

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	cmdServo = "servo"
	cmdDragTeach = "drag_teach"
	cmdGetEncode = "getEncode"
	cmdSetZeroEncode = "setZeroEncode"
	cmdStartRecord = "start_record_pos"
	cmdEndRecord = "end_record_pos"
	cmdSaveRecord = "save_record_pos"
)

func setServoStatus (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	status := vars["status"]
	
	d := &RequestDataForCommandArgs{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteBadRequestResponse(w, ERRINVALIDBODY)
		return
	}

	var params []string
	params = append(params, status)
	for _,v := range d.Args {
		params = append(params, v)
	}

	Log.Debug("servo ", status)
	SendToMCServer(w, cmdServo, params)
}

func setDragteachStatus (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	status := vars["status"]

	Log.Debug("drag_teach ", status)
	SendToMCServer(w, cmdDragTeach, ConcatParams(status))
}

func syncRobot (w http.ResponseWriter, r *http.Request){
	Log.Debug("getEncode")
	SendToMCServer(w, cmdGetEncode, nil)
}

func setZeroEncode (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	axisno := vars["axisno"]

	Log.Debug("setZeroEncode ", axisno)
	SendToMCServer(w, cmdSetZeroEncode, ConcatParams(axisno))
}

func doStartRecord (w http.ResponseWriter, r *http.Request){
    Log.Debug("start_record_pos ")
    SendToMCServer(w, cmdStartRecord, nil)
}

func doEndRecord (w http.ResponseWriter, r *http.Request){
    Log.Debug("end_record_pos ")
    SendToMCServer(w, cmdEndRecord, nil)
}

func doSaveRecord (w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    name := vars["name"]

    Log.Debug("save_record_pos ", name)
    SendToMCServer(w, cmdSaveRecord, ConcatParams(name))
}