package apiv1

import (
	"net/http"
	"encoding/json"
	"path"
	"io/ioutil"
	"os"
    "errors"

	Log "elibot-apiserver/log"

	"github.com/gorilla/mux"
)

var DBname string
var BackupPath string

func SetupDB(dbname string, backuppath string) {
	DBname = dbname
    BackupPath = backuppath
}

func DBBackupDB(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting DB Backup db")
	params := make(map[string]interface{})
	params["db_dir"] = path.Dir(DBname)

	SendToParamServer(w, "manager_backup_db", params)
}

func DBList() ([]string, error) {
    Log.Debug("in DBList")

    files, err := ioutil.ReadDir(BackupPath)
    if err != nil {
        return nil, err
    }

    var list []string
    for _, f := range files {
        list = append(list, f.Name())
    }
    return list, nil
}

func DBListBackups(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting DB List Backups")
	files, err := DBList()
	if err!=nil {
		Log.Error("fail to check list")
		WriteInternalServerErrorResponse(w, ERRLISTDBS)
		return
	}

	res, _ := json.Marshal(files)
	WriteJsonSuccessResponse(w, res)
}

func DBDel(Name string) error {
    Log.Debug("in DBDel")
    if Name == "" {
        return errors.New("file name is empty")
    }

    filename := BackupPath + Name
    if _, err := os.Stat(filename); os.IsNotExist(err) {
       return errors.New("file does not exist")
    }

    return os.Remove(filename)
}

func DBDelBackup(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting DB Backup Deletion")
	vars := mux.Vars(r)
	err := DBDel(vars["name"])
	if err!=nil {
		Log.Error("fail to delete db")
		WriteInternalServerErrorResponse(w, ERRDELETEDB)
		return
	}

	WriteSuccessResponse(w, "db deleted")
}

func DBRestoreBackup(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting DB Backup Restore")
	vars := mux.Vars(r)

	params := make(map[string]interface{})
	params["db_dir"] = path.Dir(DBname)
	params["db_bak_name"] = vars["name"]
	params["force"] = int8(1)

	SendToParamServer(w, "manager_restore_db", params)
}