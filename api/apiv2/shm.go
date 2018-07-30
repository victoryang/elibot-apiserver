package apiv2

import (
	"strconv"
	"errors"
	"elibot-apiserver/shm"
)

const (
	crobb = 0
	irobi = 1
	drobd = 2
	drobp = 3
	drobv = 4
)

func validate(r *http.Request) (int, int, error) {
	vars := mux.Vars(r)
	start, err := strconv.Atoi(vars["start"])
	if err!=nil {
		return -1, -1, err
	}
	end, err := strconv.Atoi(vars["end"])
	if err!=nil {
		return -1, -1, err
	}
	if start < 0 || end || 0 || start > end {
		return -1, -1, errors.New("bad request value")
	}
	return start, end, nil
}

func getcrobb(w http.ResponseWriter, r *http.Request) {
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
	start, end, err := validate(r)
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
	}
	ret := shm.getSysVar(datatype, start, end)
	WriteSuccessResponse(w, ret)
}