package apiv1

import (
	"net/http"
	"encoding/json"

	db "elibot-apiserver/dbproxy"
	Log "elibot-apiserver/log"

	"github.com/gorilla/mux"
)

func DBBackup(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting DB Backup")
	err := db.DBBackup()
	if err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, "succeed in backup\n")
}

func DBList(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting DB List")
	files, err := db.DBList()
	if err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	res, _ := json.Marshal(files)
	WriteSuccessResponse(w, string(res))
}

func DBDel(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting DB Deletion")
	vars := mux.Vars(r)
	err := db.DBDel(vars["name"])
	if err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, "db detelted\n")
}

func DBRestore(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting DB Restore")
	vars := mux.Vars(r)
	err := db.DBRestore(vars["name"])
	if err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, "succeed in restore\n")
}