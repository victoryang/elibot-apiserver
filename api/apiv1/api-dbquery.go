package apiv1

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

func getAllArc(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Arc")
	SendToParamServer(w, "arc_get_all", nil)
}

func getArcParams(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get Arc parameters")
	vars := mux.Vars(r)
	params := make(map[string]interface{})
	file_no, _ := strconv.Atoi(vars["file_no"])
	params["file_no"] = int32(file_no)
	params["group"] = vars["group"]

	SendToParamServer(w, "arc_get_params", params)
}

func getAllBookprograms(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all bookprograms")
	SendToParamServer(w, "bookprogram_get_all", nil)
}

func getAllEnum(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Enum")
	SendToParamServer(w, "enum_get_all", nil)
}

func getAllExtaxis(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Extaxis")
	SendToParamServer(w, "extaxis_get_all", nil)
}

func getAllInterference(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Interference")
	SendToParamServer(w, "interference_get_all", nil)
}

func getAllIos(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all IOs")
	vars := mux.Vars(r)
	params := make(map[string]interface{})
	params["group"] = vars["group"]
	params["lang"] = vars["lang"]
	auth, _ := strconv.Atoi(vars["auth"])
	params["auth"] = int32(auth)
	tech, _ := strconv.Atoi(vars["tech"])
	params["tech"] = int32(tech)

	SendToParamServer(w, "ios_get_all", params)
}

func getAllMetadata(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Metadata")
	vars := mux.Vars(r)
	params := make(map[string]interface{})
	params["lang"] = vars["lang"]

	SendToParamServer(w, "metadata_get_all", params)
}

func getAllOperationLog(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Operation Log")

	vars := mux.Vars(r)
	params := make(map[string]interface{})
	created_time, _ := strconv.Atoi(vars["created_time"])
	params["created_time"] = int32(created_time)
	start, _ := strconv.Atoi(vars["start"])
	params["start"] = int32(start)
	pageSize, _ := strconv.Atoi(vars["pageSize"])
	params["pageSize"] = int32(pageSize)

	SendToParamServer(w, "operation_record_get_all", params)
}

func getParams(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Parameter")
	SendToParamServer(w, "params_get_params", nil)
}

func getParameterById(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get Parameter by id")
	vars := mux.Vars(r)
	params := make(map[string]interface{})
	params["md_id"] = vars["md_id"]

	SendToParamServer(w, "params_get_valid_param_by_id", params)
}

func getParameterByGroup(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get Parameter by group")
	vars := mux.Vars(r)
	params := make(map[string]interface{})
	params["group"] = vars["group"]
	
	SendToParamServer(w, "params_get_valid_param_by_group", params)
}

func getAllRef(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Ref")
	SendToParamServer(w, "ref_get_all", nil)
}

func getAllToolframe(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all toolframe")
	SendToParamServer(w, "toolframe_get_all", nil)
}

func getToolframeByToolNo(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get toolframe by tool no")
	vars := mux.Vars(r)
	params := make(map[string]interface{})
	toolno, _ := strconv.Atoi(vars["toolno"])
	params["tool_no"] = int32(toolno)
	
	SendToParamServer(w, "toolframe_get_by_toolno", params)
}

func getAllUserframe(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all userframe")
	SendToParamServer(w, "userframe_get_all", nil)
}

func getUserframeByUserNo(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get userframe by user no")
	vars := mux.Vars(r)
	params := make(map[string]interface{})
	userno, _ := strconv.Atoi(vars["userno"])
	params["user_no"] = int32(userno)
	
	SendToParamServer(w, "userframe_get_by_userno", params)
}

func getAllZeroPoint(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all zeropoint")
	SendToParamServer(w, "zeropoint_get_all", nil)
}