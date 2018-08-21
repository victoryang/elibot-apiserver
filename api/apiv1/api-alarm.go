package apiv1

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"elibot-apiserver/alarm"
	Log "elibot-apiserver/log"
)

func getLogNumber(w http.ResponseWriter, r *http.Request) {
	Log.Debug("getLogNumber")
	WriteSuccessResponse(w, alarm.GetRecordsNumber())
}

func getAllLogs(w http.ResponseWriter, r *http.Request) {
	Log.Debug("getAllLogs")
	res, err := alarm.GetAllRecords()
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}

func getLogs(w http.ResponseWriter, r *http.Request) {
	Log.Debug("getLogs")
	vars := mux.Vars(r)
	from,_ := strconv.Atoi(vars["from"])
	end,_ := strconv.Atoi(vars["end"])
	if from < 0 || end < 0 || from > end {
		WriteInternalServerErrorResponse(w, ERRINCORRECTRANGE)
		return
	}
	res, err := alarm.GetRecords(from, end)
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}