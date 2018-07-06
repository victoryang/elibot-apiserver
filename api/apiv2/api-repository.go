package apiv2

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	Log "elibot-apiserver/log"
)

const (
	Tag = "apiv2:robot:repository"
)

func setArc(w http.ResponseWriter, r *http.Request) {

}

type Param struct {
	Value 		int 		`json:"value,omitempty"`
}

func setParam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	md_id := vars["md_id"]
	index := vars["index"]

	p := &Param{}
	if err := ParseBodyToJson(r, &p); err!=nil {
		WriteInternalServerErrorResponse(w, err)
		return
	}
	value := strconv.Itoa(p.Value)

	Log.Debug("set param ", md_id, value, index)
	cmd := ConcatCommand("setParam", md_id, value, index)
	SendToMCServerWithTimeout(w, r, cmd, Tag)
}