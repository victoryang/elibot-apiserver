package apiv1

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
	WriteJsonSuccessResponse(w, r)
}

type SuccessResponse struct {
	Msg 		interface{}			`json:"msg"`
}

func ToSuccessJson(msg interface{}) []byte {
	response := SuccessResponse{
		Msg:	msg,
	}
	r, _ := json.Marshal(response)
	return r
}

func WriteMultiSuccessResponse(w http.ResponseWriter, res interface{}) {
	r := ToSuccessJson(res)
	WriteJsonSuccessResponse(w, r)
}

func WriteJsonSuccessResponse(w http.ResponseWriter, res []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(res)
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
}