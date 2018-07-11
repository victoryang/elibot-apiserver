package apiv2

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	TagRepository = "apiv2:robot:repository"

	cmdSetArcParam = "setArcParam"
	cmdSetInterference = "setInterference"
	cmdSetParam = "setParam"
	cmdSetToolFrame = "setToolFrame"
	cmdSetUserFrame = "setUserFrame"
	cmdSetZeroPoint = "setZeroPoint"
)

func setArcParam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	file_no := vars["file_no"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	Log.Debug("setArcParam ", md_id, d.Value, file_no, d.Index)
	cmd := ConcatCommand(cmdSetArcParam, md_id, d.Value, file_no, d.Index)
	SendToMCServerWithTimeout(w, r, cmd, TagRepository)
}

func setInterference(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	no := vars["no"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	Log.Debug("setInterference ", md_id, d.Value, no, d.Index)
	cmd := ConcatCommand(cmdSetInterference, md_id, d.Value, no, d.Index)
	SendToMCServerWithTimeout(w, r, cmd, TagRepository)
}

func setParam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	Log.Debug("setParam ", md_id, d.Value, d.Index)
	cmd := ConcatCommand(cmdSetParam, md_id, d.Value, d.Index)
	SendToMCServerWithTimeout(w, r, cmd, TagRepository)
}

func setToolFrame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	tool_no := vars["tool_no"]
	pos_no := vars["pos_no"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	Log.Debug("setToolFrame ", md_id, d.Value, tool_no, pos_no, d.Index)
	cmd := ConcatCommand(cmdSetToolFrame, md_id, d.Value, tool_no, pos_no, d.Index)
	SendToMCServerWithTimeout(w, r, cmd, TagRepository)
}

func setUserFrame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	userno := vars["userno"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	Log.Debug("setUserFrame ", md_id, d.Value, userno)
	cmd := ConcatCommand(cmdSetUserFrame, md_id, d.Value, userno)
	SendToMCServerWithTimeout(w, r, cmd, TagRepository)
}

func setZeroPoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	Log.Debug("setZeroPoint ", md_id, d.Value, d.Index)
	cmd := ConcatCommand(cmdSetZeroPoint, md_id, d.Value, d.Index)
	SendToMCServerWithTimeout(w, r, cmd, TagRepository)
}