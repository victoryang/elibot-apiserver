package apiv1

import (
	"net/http"
	"strconv"
	"errors"

	"github.com/gorilla/mux"
)

const (
	crobb = 0
	irobi = 1
	drobd = 2
	drobp = 3
	drobv = 4

	croblb = crobb
	irobli = irobi
	drobld = drobd
	droblp = drobp
	droblv = drobv
)

func validateRange(r *http.Request) (int, int, error) {
	vars := mux.Vars(r)
	start, err := strconv.Atoi(vars["start"])
	if err!=nil {
		return -1, -1, err
	}
	end, err := strconv.Atoi(vars["end"])
	if err!=nil {
		return -1, -1, err
	}
	if start < 0 || end < 0 || start > end {
		return -1, -1, errors.New("bad request value")
	}
	return start, end, nil
}

func getcRobB(w http.ResponseWriter, r *http.Request) {
	getSysFromShm(crobb, w, r)
}

func getiRobI(w http.ResponseWriter, r *http.Request) {
	getSysFromShm(irobi, w, r)
}

func getdRobD(w http.ResponseWriter, r *http.Request) {
	getSysFromShm(drobd, w, r)
}

func getdRobP(w http.ResponseWriter, r *http.Request) {
	getSysFromShm(drobp, w, r)
}

func getdRobV(w http.ResponseWriter, r *http.Request) {
	getSysFromShm(drobv, w, r)
}

func getSysFromShm(datatype int, w http.ResponseWriter, r *http.Request) {
	start, end, err := validateRange(r)
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRINCORRECTRANGE)
		return
	}

	params := make(map[string]interface{})
	params["datatype"] = int32(datatype)
	params["start"] = int32(start)
	params["end"] = int32(end)
	SendToParamServer(w, "get_system_variables", params)
}

func getcRobLB(w http.ResponseWriter, r *http.Request) {
	getLocFromShm(croblb, w, r)
}

func getiRobLI(w http.ResponseWriter, r *http.Request) {
	getLocFromShm(irobli, w, r)
}

func getdRobLD(w http.ResponseWriter, r *http.Request) {
	getLocFromShm(drobld, w, r)
}

func getdRobLP(w http.ResponseWriter, r *http.Request) {
	getLocFromShm(droblp, w, r)
}

func getdRobLV(w http.ResponseWriter, r *http.Request) {
	getLocFromShm(droblv, w, r)
}

func getLocFromShm(datatype int32, w http.ResponseWriter, r *http.Request) {
	start, end, err := validateRange(r)
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRINCORRECTRANGE)
		return
	}
	vars := mux.Vars(r)
	num, err := strconv.Atoi(vars["num"])
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRINCORRECTRANGE)
		return
	}

	params := make(map[string]interface{})
	params["datatype"] = int32(datatype)
	params["num"] = int32(num)
	params["start"] = int32(start)
	params["end"] = int32(end)
	SendToParamServer(w, "get_local_variables", params)
}

func getPLCOnce(w http.ResponseWriter, r *http.Request) {
	SendToParamServer(w, "get_plc_once", nil)
}

func getSharedOnce(w http.ResponseWriter, r *http.Request) {
	SendToParamServer(w, "get_shared_once", nil)
}

func getNVOnce(w http.ResponseWriter, r *http.Request) {
	SendToParamServer(w, "get_nv_once", nil)
}