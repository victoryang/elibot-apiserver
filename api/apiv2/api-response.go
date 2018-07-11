package apiv2

import (
	"net/http"
	"encoding/json"
)

// Real Response
type Response struct {
	Msg 		string			`json:"msg"`
}

func ToJson(msg string) []byte {
	response := Response{
		Msg:	msg,
	}
	r, _ := json.Marshal(response)
	return r
}

// WriteSuccessResponse write success headers and response if any.
func WriteSuccessResponse(w http.ResponseWriter, res string) {
	r := ToJson(res)
	w.WriteHeader(http.StatusOK)
	w.Write(r)
}

func WriteBadRequestResponse(w http.ResponseWriter, err error) {
	r := ToJson(err.Error())
	w.WriteHeader(http.StatusBadRequest)
	w.Write(r)
}

func WriteInternalServerErrorResponse(w http.ResponseWriter, err error) {
	r := ToJson(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(r)
}