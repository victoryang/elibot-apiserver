package apiv1

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
	db "api-server/sqlitedb"
)

func getAllArc(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Arc")
	WriteJsonSuccessResponse(w, db.GetAllArc())
}

func getArcParams(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get Arc parameters")
	vars := mux.Vars(r)
	file_no, err := strconv.Atoi(vars["file_no"])
	if err!=nil {
		WriteBadRequestResponse(w, ERR_REQ_INVALID_PARAMETER)
		return
	}

	WriteJsonSuccessResponse(w, db.GetArcParams(int32(file_no), vars["group"]))
}

func getAllBookprograms(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all bookprograms")
	WriteJsonSuccessResponse(w, db.GetAllBookprograms())
}

func getAllDynamics(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all dynamics")
	WriteJsonSuccessResponse(w, db.GetAllDynamics())
}

func getDynamicsById(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get dynamics by id")
	vars := mux.Vars(r)

	WriteJsonSuccessResponse(w, db.GetDynamicsById(vars["id"]))
}

func getAllEnum(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Enum")
	WriteJsonSuccessResponse(w, db.GetAllEnum())
}

func getAllExtaxis(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Extaxis")
	WriteJsonSuccessResponse(w, db.GetAllExtaxis())
}

func getAllInterference(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Interference")
	WriteJsonSuccessResponse(w, db.GetAllInterference())
}

func getAllIos(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all IOs")
	vars := mux.Vars(r)
	auth, err := strconv.Atoi(vars["auth"])
	if err!=nil {
		WriteBadRequestResponse(w, ERR_REQ_INVALID_PARAMETER)
		return
	}

	tech, err1 := strconv.Atoi(vars["tech"])
	if err1!=nil {
		WriteBadRequestResponse(w, ERR_REQ_INVALID_PARAMETER)
		return
	}

	WriteJsonSuccessResponse(w, db.GetAllIO(vars["group"], vars["lang"], int32(auth), int32(tech)))
}

func getAllMetadata(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Metadata")
	vars := mux.Vars(r)

	WriteJsonSuccessResponse(w, db.GetAllMetadata(vars["lang"]))
}

func getAllOperationLog(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Operation Log")

	vars := mux.Vars(r)
	created_time, err := strconv.Atoi(vars["created_time"])
	if err!=nil {
		WriteBadRequestResponse(w, ERR_REQ_INVALID_PARAMETER)
		return
	}

	start, err1 := strconv.Atoi(vars["start"])
	if err1!=nil {
		WriteBadRequestResponse(w, ERR_REQ_INVALID_PARAMETER)
		return
	}

	page_size, err2 := strconv.Atoi(vars["pageSize"])
	if err2!=nil {
		WriteBadRequestResponse(w, ERR_REQ_INVALID_PARAMETER)
		return
	}

	WriteJsonSuccessResponse(w, db.GetAllOperationRecord(int32(created_time), int32(start), int32(page_size)))
}

func getParams(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Parameter")
	WriteJsonSuccessResponse(w, db.GetParams())
}

func getParameterById(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get Parameter by id")
	vars := mux.Vars(r)

	WriteJsonSuccessResponse(w, db.GetParameterById(vars["md_id"]))
}

func getParameterByGroup(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get Parameter by group")
	vars := mux.Vars(r)
	
	WriteJsonSuccessResponse(w, db.GetParameterByGroup(vars["group"]))
}

func getAllRef(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Ref")
	WriteJsonSuccessResponse(w, db.GetAllRef())
}

func getAllToolframe(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all toolframe")
	WriteJsonSuccessResponse(w, db.GetAllToolframe())
}

func getToolframeByToolNo(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get toolframe by tool no")
	vars := mux.Vars(r)
	toolno, err := strconv.Atoi(vars["toolno"])
	if err!=nil {
		WriteBadRequestResponse(w, ERR_REQ_INVALID_PARAMETER)
		return
	}
	
	WriteJsonSuccessResponse(w, db.GetToolframeByToolNo(int32(toolno)))
}

func getAllUserframe(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all userframe")
	WriteJsonSuccessResponse(w, db.GetAllUserframe())
}

func getUserframeByUserNo(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get userframe by user no")
	vars := mux.Vars(r)
	userno, err := strconv.Atoi(vars["userno"])
	if err!=nil {
		WriteBadRequestResponse(w, ERR_REQ_INVALID_PARAMETER)
		return
	}
	
	WriteJsonSuccessResponse(w, db.GetUserframeByUserNo(int32(userno)))
}

func getAllZeroPoint(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all zeropoint")
	WriteJsonSuccessResponse(w, db.GetAllZeropoint())
}