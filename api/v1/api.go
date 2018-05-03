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

func RegisterV1(r *mux.Router) http.Handler {
	r.HandleFunc("/", hello).Methods("GET")
	r.HandleFunc("/v1/test", test).Methods("GET")
	r.HandleFunc("/v1/bookprograms", getAllBookprograms).Methods("GET")
	r.HandleFunc("/v1/zeropoints", getAllZeroPoints).Methods("GET")
	r.HandleFunc("/v1/interference", getAllInterference).Methods("GET")
	return r
}