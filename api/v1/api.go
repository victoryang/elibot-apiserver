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
	fmt.Fprintf(w, "test\n")
	return
}

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "echo\n")
	return
}

func RegisterV1(r *mux.Router) http.Handler {
	r.HandleFunc("/", hello).Methods("GET")
	r.HandleFunc("/api/v1/test", test).Methods("GET")
	return r
}