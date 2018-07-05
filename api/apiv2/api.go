package apiv2

import (
	"net/http"
	"io/ioutil"
	"errors"
	"context"
	"time"
	"encoding/json"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
	"elibot-apiserver/mcserver"
)

const (
	CommandEnd = "\n"
	Space = " "
)

type Body struct {
	Data 		int 		`json:"data,omitempty"`
}

func setParam(w http.ResponseWriter, r *http.Request) {
	var mcs *mcserver.MCserver
	if mcs = mcserver.GetMcServer(); mcs == nil {
		WriteInternalServerErrorResponse(w, errors.New("mcserver is not available right now"))
		return
	}

	vars := mux.Vars(r)
	md_id := vars["md_id"] + Space
	index := vars["index"] + CommandEnd

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}
	var b Body
	if err := json.Unmarshal(body, &b); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}
	value := strconv.Itoa(b.Data) + Space

	cmd := "setParam " + md_id + value + index
	tag := "restapi:apiv2"
	rCh := make(chan mcserver.Response)
	defer close(rCh)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	go mcs.Exec(cmd, tag, rCh)

	select {
	case <-ctx.Done():
		WriteInternalServerErrorResponse(w, ctx.Err())
	case r := <- rCh:
		if r.Err != nil {
			WriteInternalServerErrorResponse(w, r.Err)
		} else {
			WriteSuccessResponse(w, r.Result)
		}
	}
}