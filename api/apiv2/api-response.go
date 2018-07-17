package apiv2

import (
	"net/http"
	"encoding/json"
)

/* respond differently*/
type SuccessResponse struct {
	Msg 		string			`json:"msg"`
}

func ToSuccessJson(msg string) []byte {
	response := SuccessResponse{
		Msg:	msg,
	}
	r, _ := json.Marshal(response)
	return r
}

// WriteSuccessResponse write success headers and response if any.
func WriteSuccessResponse(w http.ResponseWriter, res string) {
	r := ToSuccessJson(res)
	w.WriteHeader(http.StatusOK)
	w.Write(r)
}

type ErrorResponse struct {
	ErrCode		int 			`json:"error_code"`
	ErrMsg 		string			`json:"error_msg"`
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
}