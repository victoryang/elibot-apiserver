package v1

import (
	"net/http"
	"fmt"
	"errors"

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
	rCh := make(chan mcserver.Response)
	mcs.Exec(cmd, from, rCh)
	rr := <- rCh
	res := rr.Result
	err := rr.Err
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}
	
	WriteSuccessResponse(w, res)
}