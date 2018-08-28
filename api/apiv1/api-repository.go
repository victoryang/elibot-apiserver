package apiv1

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
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
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("setArcParam ", md_id, d.Value, file_no, strconv.Itoa(d.Index))
	SendToMCServerWithJsonRpc(w, cmdSetArcParam, ConcatParams(md_id, d.Value, file_no, strconv.Itoa(d.Index)))
}

func setInterference(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	no := vars["no"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("setInterference ", md_id, d.Value, no, strconv.Itoa(d.Index))
	SendToMCServerWithJsonRpc(w, cmdSetInterference, ConcatParams(md_id, d.Value, no, strconv.Itoa(d.Index)))
}

func setParam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("setParam ", md_id, d.Value, strconv.Itoa(d.Index))
	SendToMCServerWithJsonRpc(w, cmdSetParam, ConcatParams(md_id, d.Value, strconv.Itoa(d.Index)))
}

func setToolFrame(w http.ResponseWriter, r *http.Request) {
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
	SendToMCServerWithJsonRpc(w, cmdSetToolFrame, ConcatParams(md_id, d.Value, tool_no, pos_no, strconv.Itoa(d.Index)))
}

func setUserFrame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	userno := vars["userno"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("setUserFrame ", md_id, d.Value, userno)
	SendToMCServerWithJsonRpc(w, cmdSetUserFrame, ConcatParams(md_id, d.Value, userno))
}

func setZeroPoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]

	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}

	Log.Debug("setZeroPoint ", md_id, d.Value, strconv.Itoa(d.Index))
	SendToMCServerWithJsonRpc(w, cmdSetZeroPoint, ConcatParams(md_id, d.Value, strconv.Itoa(d.Index)))
}