package apiv1

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"elibot-apiserver/alarm"
	Log "elibot-apiserver/log"
)

func getLogs(w http.ResponseWriter, r *http.Request) {
	Log.Debug("getLogs")
	vars := mux.Vars(r)
	start,_ := strconv.Atoi(vars["start"])
	end,_ := strconv.Atoi(vars["end"])
	if start < 0 || end < 0 || start > end {
		WriteBadRequestResponse(w, ERRINCORRECTRANGE)
		return
	}
	timestamp,_ := strconv.Atoi(vars["timestamp"])
	res, err := alarm.GetRecords(start, end, uint32(timestamp))
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getLogsByAlarmLevel(w http.ResponseWriter, r *http.Request) {
	Log.Debug("getLogsByAlarmLevel")
	vars := mux.Vars(r)
	start,_ := strconv.Atoi(vars["start"])
	end,_ := strconv.Atoi(vars["end"])
	if start < 0 || end < 0 || start > end {
		WriteBadRequestResponse(w, ERRINCORRECTRANGE)
		return
	}

	timestamp,_ := strconv.Atoi(vars["timestamp"])
	res, err := alarm.GetRecordsByLevel(vars["level"], start, end, uint32(timestamp))
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}