package v1

import (
	"net/http"

	Log "elibot-apiserver/log"

	"github.com/gorilla/mux"
)

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

func RegisterAPIRouter(r *mux.Router) http.Handler {
	Log.Debug("Register V1 api...")
	return RegisterV1(r)
}