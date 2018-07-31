package apiv1

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	db "elibot-apiserver/dbproxy"
	Log "elibot-apiserver/log"
)

func getAllArc(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Arc")
	res, err := db.Get_ALL_Arc()
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteSuccessResponse(w, res)
}

func getArcParams(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get Arc parameters")
	vars := mux.Vars(r)
	queries := make(map[string]interface{})

	file_no, _ := strconv.Atoi(vars["file_no"])
	queries["file_no"] = int32(file_no)
	queries["group"] = vars["group"]
	res, err := db.Get_Arc_Params(queries)
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllBookprograms(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all bookprograms")
	res, err := db.Get_All_Bookprograms()
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllEnum(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Enum")
	res, err := db.Get_ALL_Enum()
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllExtaxis(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Extaxis")
	res, err := db.Get_ALL_Extaxis()
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllInterference(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Interference")
	res, err := db.Get_All_Interference()
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllIO(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all IO")
	res, err := db.Get_All_IO()
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllMetadata(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Metadata")
	vars := mux.Vars(r)

	queries := make(map[string]interface{})
	queries["lang"] = vars["lang"]
	res, err := db.Get_All_Metadata(queries)
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteSuccessResponse(w, res)
}

func getParams(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Parameter")
	res, err := db.Get_Params()
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteSuccessResponse(w, res)
}

func getParameterById(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Parameter")
	vars := mux.Vars(r)
	queries := make(map[string]interface{})
	queries["md_id"] = vars["md_id"]

	res, err := db.Get_Parameter_By_Id(queries)
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteSuccessResponse(w, res)
}

func parameterbygroup(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Parameter")
	vars := mux.Vars(r)
	queries := make(map[string]interface{})
	queries["group"] = vars["group"]
	res, err := db.Get_Parameter_By_Group(queries)
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllRef(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Ref")
	res, err := db.Get_All_Ref()
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllToolframe(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all toolframe")
	res, err := db.Get_ALL_Toolframe()
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllUserframe(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all userframe")
	res, err := db.Get_ALL_Userframe()
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllZeroPoints(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all zeropoints")
	res, err := db.Get_All_Zeropoints()
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteSuccessResponse(w, res)
}