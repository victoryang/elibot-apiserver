package apiv1

import (
	"net/http"
	"encoding/json"
)

// WriteSuccessResponse write success headers and response if any.
func WriteSuccessResponse(w http.ResponseWriter, response []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

type ErrorResponse struct {
	ErrCode		int 			`json:"errcode"`
	ErrMsg 		string			`json:"errmsg"`
}

func ToErrJson(errno int) []byte {
	response := ErrorResponse{
		ErrCode:	errno,
		ErrMsg:		ErrMsg(errno),
	}
	r, _ := json.Marshal(response)
	return r
}

func WriteInternalServerErrorResponse(w http.ResponseWriter, errno int) {
	r := ToErrJson(errno)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(r)
	w.Flush()
}