package apiv2

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	Tag = "apiv2:robot:repository"

	cmdSetArcParam = "setArcParam"
	cmdSetInterference = "setInterference"
	cmdSetParam = "setParam"
	cmdSetToolFrame = "setToolFrame"
	cmdSetUserFrame = "setUserFrame"
	cmdSetZeroPoint = "setZeroPoint"
)

type Default struct {
	Value 		string 		`json:"value,omitempty"`
}

func setArcParam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	file_no := vars["file_no"]
	index := vars["index"]

	d := &Default{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	Log.Debug("setArcParam ", md_id, d.Value, index)
	cmd := ConcatCommand(cmdSetArcParam, md_id, d.Value, file_no, index)
	SendToMCServerWithTimeout(w, r, cmd, Tag)
}

func setInterference(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	no := vars["no"]
	index := vars["index"]

	d := &Default{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	Log.Debug("setInterference ", md_id, d.Value, no, index)
	cmd := ConcatCommand(cmdSetInterference, md_id, d.Value, no, index)
	SendToMCServerWithTimeout(w, r, cmd, Tag)
}

func setParam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	index := vars["index"]

	d := &Default{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	Log.Debug("setParam ", md_id, d.Value, index)
	cmd := ConcatCommand(cmdSetParam, md_id, d.Value, index)
	SendToMCServerWithTimeout(w, r, cmd, Tag)
}

func setToolFrame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	index := vars["index"]
	index1 := vars["index1"]
	index2 := vars["index2"]

	d := &Default{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	tmp := make([]string, 0)
	tmp = append(tmp, md_id)
	Log.Debug("setToolFrame ", md_id, d.Value, index, index1, index2)
	cmd := ConcatCommand(cmdSetToolFrame, md_id, d.Value, index, index1, index2)
	SendToMCServerWithTimeout(w, r, cmd, Tag)
}

func setUserFrame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	userNo := vars["userNo"]

	d := &Default{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	tmp := make([]string, 0)
	tmp = append(tmp, md_id)
	Log.Debug("setUserFrame ", md_id, d.Value, userNo)
	cmd := ConcatCommand(cmdSetUserFrame, md_id, d.Value, userNo)
	SendToMCServerWithTimeout(w, r, cmd, Tag)
}

func setZeroPoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	index := vars["index"]

	d := &Default{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	Log.Debug("setZeroPoint ", md_id, d.Value, index)
	cmd := ConcatCommand(cmdSetZeroPoint, md_id, d.Value, index)
	SendToMCServerWithTimeout(w, r, cmd, Tag)
}