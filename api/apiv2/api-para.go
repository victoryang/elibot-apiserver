package apiv2

import (
	"net/http"

	//"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	TagParameter = "apiv2:robot:parameter"

	cmdSetParameters = "para"
)

func setParameters(w http.ResponseWriter, r *http.Request) {
	d := &RequestData{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	Log.Debug("para ", d.Index, d.Value)
	cmd := ConcatCommand(cmdSetParameters, d.Index, d.Value)
	SendToMCServerWithTimeout(w, r, cmd, TagParameter)
}