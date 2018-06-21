package v1

import (
	"net/http"
	"fmt"
	"errors"
	"strconv"
	"encoding/json"

	"github.com/gorilla/mux"

	db "elibot-apiserver/dbproxy"
	Log "elibot-apiserver/log"
	"elibot-apiserver/mcserver"
)

func hello(w http.ResponseWriter, r *http.Request) {
	WriteSuccessResponse(w, "Welcome to elibot\n")
}

func test(w http.ResponseWriter, r *http.Request) {
	WriteSuccessResponse(w, "echo\n")
}

func testSocket(w http.ResponseWriter, r *http.Request) {
	var mcs *mcserver.MCserver
	if mcs = mcserver.GetMcServer(); mcs == nil {
		WriteInternalServerErrorResponse(w, errors.New("mcserver is not available right now"))
		return
	}
	cmd := "testGo 0 1\n"
	from := "restapi:testsocket"
	resp := make(chan mcserver.Response)
	go mcs.Exec(cmd, from, resp)
	rr := <-resp
	res := rr.Result
	err := rr.Err
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}
	
	WriteSuccessResponse(w, res)
}

func getAllArc(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Arc")
	res, err := db.Get_ALL_Arc()
	if err!=nil {
		WriteInternalServerErrorResponse(w, err)
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
		WriteInternalServerErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllBookprograms(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all bookprograms")
	res, err := db.Get_All_Bookprograms()
	if err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllEnum(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Enum")
	res, err := db.Get_ALL_Enum()
	if err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllExtaxis(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Extaxis")
	res, err := db.Get_ALL_Extaxis()
	if err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllInterference(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Interference")
	res, err := db.Get_All_Interference()
	if err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllIO(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all IO")
	res, err := db.Get_All_IO()
	if err!=nil {
		WriteInternalServerErrorResponse(w, err)
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
		WriteInternalServerErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, res)
}

func getParams(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Parameter")
	res, err := db.Get_Params()
	if err!=nil {
		WriteInternalServerErrorResponse(w, err)
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
		WriteInternalServerErrorResponse(w, err)
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
		WriteInternalServerErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllRef(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all Ref")
	res, err := db.Get_All_Ref()
	if err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllToolframe(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all toolframe")
	res, err := db.Get_ALL_Toolframe()
	if err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllUserframe(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all userframe")
	res, err := db.Get_ALL_Userframe()
	if err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, res)
}

func getAllZeroPoints(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get all zeropoints")
	res, err := db.Get_All_Zeropoints()
	if err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	WriteSuccessResponse(w, res)
}

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

func RegisterV1(r *mux.Router) http.Handler {
	r.HandleFunc("/", hello).Methods("GET")
	r.HandleFunc("/v1/test", test).Methods("GET")
	r.HandleFunc("/v1/testSocket", testSocket).Methods("GET")

	r.HandleFunc("/v1/arc", getAllArc).Methods("GET")
	r.HandleFunc("/v1/arcparams", getArcParams).Methods("GET").Queries("file_no", "{file_no}", "group", "{group}")
	r.HandleFunc("/v1/bookprograms", getAllBookprograms).Methods("GET")
	r.HandleFunc("/v1/enum", getAllEnum).Methods("GET")
	r.HandleFunc("/v1/extaxis", getAllExtaxis).Methods("GET")
	r.HandleFunc("/v1/interference", getAllInterference).Methods("GET")
	//r.HandleFunc("/v1/io", getAllIO).Methods("GET")
	r.HandleFunc("/v1/metadata", getAllMetadata).Methods("GET").Queries("lang", "{lang}")
	r.HandleFunc("/v1/params", getParams).Methods("GET")
	r.HandleFunc("/v1/parameterbyid", getParameterById).Methods("GET").Queries("md_id", "{md_id}")
	r.HandleFunc("/v1/parameterbygroup", parameterbygroup).Methods("GET").Queries("group", "{group}")
	r.HandleFunc("/v1/ref", getAllRef).Methods("GET")
	r.HandleFunc("/v1/toolframe", getAllToolframe).Methods("GET")
	r.HandleFunc("/v1/userframe", getAllUserframe).Methods("GET")
	r.HandleFunc("/v1/zeropoints", getAllZeroPoints).Methods("GET")

	/*For DB backup, restore and upgrade*/
	r.HandleFunc("/v1/db/backup", DBBackup).Methods("POST")
	r.HandleFunc("/v1/db/backup", DBList).Methods("GET")
	r.HandleFunc("/v1/db/backup", DBDel).Methods("DELETE").Queries("name", "{name}")
	r.HandleFunc("/v1/db/restore", DBRestore).Methods("POST").Queries("name", "{name}")
	return r
}