package v1

import (
	"net/http"
	"fmt"
	"strconv"
	"encoding/json"

	"github.com/gorilla/mux"

	db "elibot-apiserver/dbproxy"
	Log "elibot-apiserver/log"
	"elibot-apiserver/mcserver"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to elibot\n")
	return
}

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "echo\n")
	return
}

func testSocket(w http.ResponseWriter, r *http.Request) {
	resp := make(chan mcserver.Response)
	mcserver.Mcs.WorkChan<-mcserver.Request{Command: "testGo 0 1\n", From:"restapi:testsocket", Resp:resp}
	rr := <-resp
	res := rr.Result
	err := rr.Err
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}
	
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllArc(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting get all Arc")
	res, err := db.Get_ALL_Arc()
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getArcParams(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting get Arc parameters")
	vars := mux.Vars(r)
	queries := make(map[string]interface{})

	file_no, _ := strconv.Atoi(vars["file_no"])
	queries["file_no"] = int32(file_no)
	queries["group"] = vars["group"]
	res, err := db.Get_Arc_Params(queries)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllBookprograms(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting get all bookprograms")
	res, err := db.Get_All_Bookprograms()
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllEnum(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting get all Enum")
	res, err := db.Get_ALL_Enum()
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllExtaxis(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting get all Extaxis")
	res, err := db.Get_ALL_Extaxis()
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllInterference(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting get all Interference")
	res, err := db.Get_All_Interference()
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllIO(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting get all IO")
	res, err := db.Get_All_IO()
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllMetadata(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting get all Metadata")
	vars := mux.Vars(r)

	queries := make(map[string]interface{})
	queries["lang"] = vars["lang"]
	res, err := db.Get_All_Metadata(queries)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getParams(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting get all Parameter")
	res, err := db.Get_Params()
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getParameterById(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting get all Parameter")
	vars := mux.Vars(r)
	queries := make(map[string]interface{})
	queries["md_id"] = vars["md_id"]

	res, err := db.Get_Parameter_By_Id(queries)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func parameterbygroup(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting get all Parameter")
	vars := mux.Vars(r)
	queries := make(map[string]interface{})
	queries["group"] = vars["group"]
	res, err := db.Get_Parameter_By_Group(queries)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllRef(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting get all Ref")
	res, err := db.Get_All_Ref()
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllToolframe(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting get all toolframe")
	res, err := db.Get_ALL_Toolframe()
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllUserframe(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting get all userframe")
	res, err := db.Get_ALL_Userframe()
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllZeroPoints(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting get all zeropoints")
	res, err := db.Get_All_Zeropoints()
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func DBBackup(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting DB Backup")
	err := db.DBBackup()
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "succeed in backup\n")
	return
}

func DBList(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting DB List")
	files, err := db.DBList()
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	res, _ := json.Marshal(files)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(res))
	return
}

func DBDel(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting DB Deletion")
	vars := mux.Vars(r)
	err := db.DBDel(vars["name"])
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "db detelted\n")
	return
}

func DBRestore(w http.ResponseWriter, r *http.Request) {
	Log.Info("starting DB Restore")
	vars := mux.Vars(r)
	err := db.DBRestore(vars["name"])
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "succeed in restore\n")
	return
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