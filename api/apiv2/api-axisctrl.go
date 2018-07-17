package apiv2

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	TagAxisCtrl = "apiv2:robot:axisctrl"

	cmdDynamicMode = "dynamicMode"
	cmdTestSlotIO = "test_slotIO"
	cmdGetEncode = "getEncode"
	cmdSetZeroEncode = "setZeroEncode"
	cmdSetServoStatus = "servo"
	cmdGetAxisInput = "get_axis_input"
	cmdSetAxisOnput = "set_axis_output"
	cmdDragTeach = "drag_teach"
)

func setDynamicMode(w http.ResponseWriter, r *http.Request) {
	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("dynamicMode ", d.Value)
	cmd := ConcatCommand(cmdDynamicMode, d.Value)
	SendToMCServerWithTimeout(w, r, cmd, TagAxisCtrl)
}

func testSlotIO (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	slotnum := vars["slotnum"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("test_slotIO ", slotnum, d.Value)
	cmd := ConcatCommand(cmdTestSlotIO, slotnum, d.Value)
	SendToMCServerWithTimeout(w, r, cmd, TagAxisCtrl)
}

func getEncode (w http.ResponseWriter, r *http.Request){
	Log.Debug("getEncode")
	SendToMCServerWithTimeout(w, r, cmdGetEncode, TagAxisCtrl)
}

func setZeroEncode (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	axis := vars["axis"]

	Log.Debug("setZeroEncode ", axis)
	cmd := ConcatCommand(cmdSetZeroEncode, axis)
	SendToMCServerWithTimeout(w, r, cmd, TagAxisCtrl)
}

func setServoStatus (w http.ResponseWriter, r *http.Request){
	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("servo ", d.Value)
	cmd := ConcatCommand(cmdSetServoStatus, d.Value)
	SendToMCServerWithTimeout(w, r, cmd, TagAxisCtrl)
}

func getAxisInput (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	channel := vars["channel"]
	io_num := vars["io_num"]

	Log.Debug("get_axis_input ", channel, io_num)
	cmd := ConcatCommand(cmdGetAxisInput, channel, io_num)
	SendToMCServerWithTimeout(w, r, cmd, TagAxisCtrl)
}

func setAxisOnput (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	channel := vars["channel"]
	io_num := vars["io_num"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("set_axis_output ", channel, io_num, d.Value)
	cmd := ConcatCommand(cmdSetAxisOnput, channel, io_num, d.Value)
	SendToMCServerWithTimeout(w, r, cmd, TagAxisCtrl)
}

func setDragteach (w http.ResponseWriter, r *http.Request){
	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("drag_teach ", d.Value)
	cmd := ConcatCommand(cmdDragTeach, d.Value)
	SendToMCServerWithTimeout(w, r, cmd, TagAxisCtrl)
}

/*func getEncoder (w http.ResponseWriter, r *http.Request){
	d := &Default{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	Log.Debug("drag_teach ", d.Value)
	cmd := ConcatCommand(cmdDragTeach, d.Value)
	SendToMCServerWithTimeout(w, r, cmd, TagAxisCtrl)
}*/