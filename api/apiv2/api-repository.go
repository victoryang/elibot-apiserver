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

// swagger:route PUT /v2/robot/repository/arcparam/{file_no}/{md_id} setArcParam
// Origin: setArcParam [md_id] [value] [file_no] [index]
// This allows to set arcparam with some parameters
//
// 		Consumes:
//		- application/json
//		- Request
//		Produces:
//		- application/json
//		Schemes: http
//		Responses:
//			default: Response
//
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

// swagger:route PUT /v2/robot/repository/interference/{no}/{md_id} setInterference
// Origin: setInterference [md_id] [value] [no] [index]
// This allows to set interference with some parameters
//
// 		Consumes:
//		- application/json
//		- Request
//		Produces:
//		- application/json
//		Schemes: http
//		Responses:
//			default: Response
//
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

// swagger:route PUT /v2/robot/repository/params/{md_id} setParam
// Origin: setParam [md_id] [value] [index]
// This allows to set param with some parameters
//
// 		Consumes:
//		- application/json
//		- Request
//		Produces:
//		- application/json
//		Schemes: http
//		Responses:
//			default: Response
//
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

// swagger:route PUT /v2/robot/repository/toolframes/{md_id} setToolFrame
// Origin: setToolFrame [md_id] [value] [tool_no] [pos_no|index] [index]
// This allows to set toolframe with some parameters
//
// 		Consumes:
//		- application/json
//		- Request
//		Produces:
//		- application/json
//		Schemes: http
//		Responses:
//			default: Response
//
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

// swagger:route PUT /v2/robot/repository/userframe/{userno}/{md_id} setUserFrame
// Origin: setUserFrame [md_id] [value] [userNo]
// This allows to set userframe with some parameters
//
// 		Consumes:
//		- application/json
//		- Request
//		Produces:
//		- application/json
//		Schemes: http
//		Responses:
//			default: Response
//
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

// swagger:route PUT /v2/robot/repository/zeropoint/{md_id} setZeroPoint
// Origin: setZeroPoint [md_id] [value] [index]
// This allows to set zeropoint with some parameters
//
// 		Consumes:
//		- application/json
//		- Request
//		Produces:
//		- application/json
//		Schemes: http
//		Responses:
//			default: Response
//
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