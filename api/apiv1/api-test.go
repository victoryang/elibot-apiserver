package apiv1

import (
	"net/http"
	"context"
	"time"
	"errors"

	"elibot-apiserver/mcserver"
)

func hello(w http.ResponseWriter, r *http.Request) {
	WriteSuccessResponse(w, []byte("Welcome to elibot\n"))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	WriteSuccessResponse(w, []byte("echo OK\n"))
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
	defer close(rCh)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	go mcs.Exec(cmd, from, rCh)

	select {
	case <-ctx.Done():
		WriteInternalServerErrorResponse(w, ctx.Err())
	case r := <- rCh:
		if r.Err != nil {
			WriteInternalServerErrorResponse(w, r.Err)
		} else {
			WriteSuccessResponse(w, []byte(r.Result))
		}
	}
}