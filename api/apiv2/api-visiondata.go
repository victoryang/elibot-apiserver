package apiv2

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	TagVisionData = "apiv2:robot:visiondata"

	cmdSetVisionCraftNum = "set_visioncraftnum"
	cmdSetVisionDataNote = "set_visiondatanote"
)

type Note struct {
	Value  		[]string		`json:"value,omitempty"`
}

func setVisionCraftNum(w http.ResponseWriter, r *http.Request) {
	d := &Default{}
	if err := ParseBodyToObject(r, d); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	Log.Debug("set_visioncraftnum ", d.Value)
	cmd := ConcatCommand(cmdSetVisionCraftNum, d.Value)
	SendToMCServerWithTimeout(w, r, cmd, TagVisionData)
}

func setVisionDataNote(w http.ResponseWriter, r *http.Request) {
	n := &Note{}
	if err := ParseBodyToObject(r, n); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}

	Log.Debug("set_visiondatanote ", n.Value)
	cmd := ConcatCommand(cmdSetVisionDataNote, n.Value...)
	SendToMCServerWithTimeout(w, r, cmd, TagVisionData)
}