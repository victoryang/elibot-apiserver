package apiv1

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"elibot-apiserver/alarm"
	Log "elibot-apiserver/log"
)

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
	from := strconv.Atoi(vars["from"])
	end := strconv.Atoi(vars["end"])
	res, err := alarm.GetRecords(from, end)
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRQUERY)
		return
	}

	WriteJsonSuccessResponse(w, res)
}