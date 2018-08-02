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
	errno := db.DBBackup()
	if errno!=0 {
		Log.Error("fail to backup db")
		WriteInternalServerErrorResponse(w, ERRBACKUPDB)
		return
	}

	WriteSuccessResponse(w, "succeed in backup")
}

func DBList(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting DB List")
	files, err := db.DBList()
	if err!=nil {
		Log.Error("fail to check list")
		WriteInternalServerErrorResponse(w, ERRLISTDBS)
		return
	}

	WriteSuccessResponse(w, files)
}

func DBDel(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting DB Deletion")
	vars := mux.Vars(r)
	err := db.DBDel(vars["name"])
	if err!=nil {
		Log.Error("fail to delete db")
		WriteInternalServerErrorResponse(w, ERRDELETEDB)
		return
	}

	WriteSuccessResponse(w, "db detelted")
}

func DBRestore(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting DB Restore")
	vars := mux.Vars(r)
	errno := db.DBRestore(vars["name"])
	if errno!=0 {
		Log.Error("fail to restore db")
		WriteInternalServerErrorResponse(w, ERRRESTOREDB)
		return
	}

	WriteSuccessResponse(w, "succeed in restore")
}