package v1

import (
	"net/http"
	"fmt"
	"encoding/json"

	"github.com/gorilla/mux"
	"elibot-apiserver/db"
)

type

func hello(w http.ResponseWriter, r *http.Request) {
	db.NewDBContext("elibotDB.db")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "test\n")
	return
}

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "echo\n")
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
	fmt.Fprintf(w, string(res))
	return
}

func RegisterV1(r *mux.Router) http.Handler {
	r.HandleFunc("/", hello).Methods("GET")
	r.HandleFunc("/api/v1/test", test).Methods("GET")
	r.HandleFunc("/test", getAllBookprograms).Methods("GET")
	return r
}