package v1

import (
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	"elibot-apiserver/db"
)

func hello(w http.ResponseWriter, r *http.Request) {
	db.NewDBContext("elibotDB.db")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to elibot\n")
	return
}

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "echo\n")
	return
}

func getAllArc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting get all Arc")
	res, err := db.Get_ALL_Arc()
	if err!=nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getArcParams(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting get Arc parameters")
	queries := make(map[string]string)
	queries["file_no"] = r.URL.Query().Get("file_no")
	res, err := db.Get_Arc_Params(queries)
	if err!=nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllBackup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting get all Backup")
	res, err := db.Get_All_Backup()
	if err!=nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllBookprograms(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting get all bookprograms")
	res, err := db.Get_All_Bookprograms()
	if err!=nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllEnum(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting get all Enum")
	res, err := db.Get_ALL_Enum()
	if err!=nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllExtaxis(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting get all Extaxis")
	res, err := db.Get_ALL_Extaxis()
	if err!=nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllInterference(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting get all Interference")
	res, err := db.Get_All_Interference()
	if err!=nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllIO(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting get all IO")
	res, err := db.Get_All_IO()
	if err!=nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllMetadata(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting get all Metadata")
	res, err := db.Get_All_Metadata()
	if err!=nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllParameter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting get all Parameter")
	res, err := db.Get_All_Parameter()
	if err!=nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllRef(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting get all Ref")
	res, err := db.Get_All_Ref()
	if err!=nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllToolframe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting get all toolframe")
	res, err := db.Get_ALL_Toolframe()
	if err!=nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllUserframe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting get all userframe")
	res, err := db.Get_ALL_Userframe()
	if err!=nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func getAllZeroPoints(w http.ResponseWriter, r *http.Request) {
	fmt.Println("starting get all zeropoints")
	res, err := db.Get_All_Zeropoints()
	if err!=nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, res)
	return
}

func RegisterV1(r *mux.Router) http.Handler {
	r.HandleFunc("/", hello).Methods("GET")
	r.HandleFunc("/v1/test", test).Methods("GET")
	r.HandleFunc("/v1/arc", getAllArc).Methods("GET")
	r.HandleFunc("/v1/arcparams", getArcParams).Methods("GET").Queries("file_no", "{file_no}")
	//r.HandleFunc("/v1/backup", getAllBackup).Methods("GET")
	r.HandleFunc("/v1/bookprograms", getAllBookprograms).Methods("GET")
	//r.HandleFunc("/v1/enum", getAllEnum).Methods("GET")
	r.HandleFunc("/v1/extaxis", getAllExtaxis).Methods("GET")
	r.HandleFunc("/v1/interference", getAllInterference).Methods("GET")
	//r.HandleFunc("/v1/io", getAllIO).Methods("GET")
	//r.HandleFunc("/v1/metadata", getAllMetadata).Methods("GET")
	r.HandleFunc("/v1/parameter", getAllParameter).Methods("GET")
	r.HandleFunc("/v1/ref", getAllRef).Methods("GET")
	r.HandleFunc("/v1/toolframe", getAllToolframe).Methods("GET")
	r.HandleFunc("/v1/userframe", getAllUserframe).Methods("GET")
	r.HandleFunc("/v1/zeropoints", getAllZeroPoints).Methods("GET")
	return r
}