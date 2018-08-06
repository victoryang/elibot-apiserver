package apiv1

import (
	"context"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"elibot-apiserver/mcserver"
	Log "elibot-apiserver/log"
)

const (
	duration = 3
	space = " "
	endline = "\n"
)
var ctx_base context.Context
var mcs *mcserver.MCserver

func init() {
	ctx_base = context.Background()
}

func ConcatCommand(cmd string, vars ...string) string {
	var command string
	command += cmd
	for _, v := range vars {
		if v == "" {
			continue
		}
		command += space
		command += v
	}
	command += endline
	return command
}

func ParseBodyToObject(r *http.Request, des interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err!=nil {
		Log.Error("Parse fails: ", err)
		return err
	}

	if err := json.Unmarshal(body, des); err!=nil {
		Log.Error("Parse fails: ", err)
		return err
	}
	return nil
}

func SendToMCServerWithTimeout(w http.ResponseWriter, r *http.Request, cmd string, tag string) {
	if mcs = mcserver.GetMcServer(); mcs == nil {
		Log.Error("mcserver is not available right now")
		WriteInternalServerErrorResponse(w, ERRMCSEVERNOTAVAILABLE)
		return
	}

	rCh := make(chan mcserver.Response)
	defer close(rCh)

	ctx, cancel := context.WithTimeout(ctx_base, duration*time.Second)
	defer cancel()
	
	go mcs.Exec(cmd, tag, rCh)

	select {
	case <-ctx.Done():
		Log.Debug("Request to mcserver times out or cancelled: ", ctx.Err())
		WriteInternalServerErrorResponse(w, ERRREQUESTTIMEOUT)
	case r := <- rCh:
		if r.Err != nil {
			Log.Debug("Request to mcserver fails, due to: ", r.Err)
			WriteInternalServerErrorResponse(w, ERRREQUESTFAIL)
		} else {
			WriteSuccessResponse(w, r.Result)
		}
	}
}