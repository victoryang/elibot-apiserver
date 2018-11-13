package apiv1

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	db "elibot-apiserver/sqlitedb"
	Log "elibot-apiserver/log"
)

func getAllArc(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Arc")
	res, err := db.GetAllArc()
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getArcParams(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get Arc parameters")
	vars := mux.Vars(r)
	queries := make(map[string]interface{})

	file_no, _ := strconv.Atoi(vars["file_no"])
	queries["file_no"] = int32(file_no)
	queries["group"] = vars["group"]
	res, err := db.GetArcParams(queries)
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getAllBookprograms(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all bookprograms")
	res, err := db.GetAllBookprograms()
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getAllEnum(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Enum")
	res, err := db.GetAllEnum()
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getAllExtaxis(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Extaxis")
	res, err := db.GetAllExtaxis()
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getAllInterference(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Interference")
	res, err := db.GetAllInterference()
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getAllIos(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all IOs")
	vars := mux.Vars(r)
	queries := make(map[string]interface{})
	queries["group"] = vars["group"]
	queries["lang"] = vars["lang"]
	auth, _ := strconv.Atoi(vars["auth"])
	queries["auth"] = int32(auth)
	tech, _ := strconv.Atoi(vars["tech"])
	queries["tech"] = int32(tech)

	res, err := db.GetAllIO(queries)
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getAllMetadata(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Metadata")
	vars := mux.Vars(r)

	queries := make(map[string]interface{})
	queries["lang"] = vars["lang"]
	res, err := db.GetAllMetadata(queries)
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getParams(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Parameter")
	res, err := db.GetParams()
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getParameterById(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get Parameter by id")
	vars := mux.Vars(r)
	queries := make(map[string]interface{})
	queries["md_id"] = vars["md_id"]

	res, err := db.GetParameterById(queries)
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getParameterByGroup(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get Parameter by group")
	vars := mux.Vars(r)
	queries := make(map[string]interface{})
	queries["group"] = vars["group"]
	res, err := db.GetParameterByGroup(queries)
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getAllRef(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Ref")
	res, err := db.GetAllRef()
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getAllToolframe(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all toolframe")
	res, err := db.GetAllToolframe()
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getToolframeByToolNo(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get toolframe by tool no")
	vars := mux.Vars(r)
	queries := make(map[string]interface{})
	toolno, _ := strconv.Atoi(vars["toolno"])
	queries["tool_no"] = int32(toolno)
	res, err := db.GetToolframeByToolNo(queries)
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getAllUserframe(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all userframe")
	res, err := db.GetAllUserframe()
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getUserframeByUserNo(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get userframe by user no")
	vars := mux.Vars(r)
	queries := make(map[string]interface{})
	userno, _ := strconv.Atoi(vars["userno"])
	queries["user_no"] = int32(userno)
	res, err := db.GetUserframeByUserNo(queries)
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getAllZeroPoint(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all zeropoint")
	res, err := db.GetAllZeropoint()
	if err!=nil {
		Log.Error("Fail to query: ", err)
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}