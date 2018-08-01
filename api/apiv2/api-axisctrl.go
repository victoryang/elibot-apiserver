package apiv2

import (
	"net/http"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	TagAxisCtrl = "apiv2:robot:axisctrl"

	cmdDragTeach = "drag_teach"
)

func setDragteach (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	enabled := vars["enabled"]

	Log.Debug("drag_teach ", enabled)
	cmd := ConcatCommand(cmdDragTeach, enabled)
	SendToMCServerWithTimeout(w, r, cmd, TagAxisCtrl)
}