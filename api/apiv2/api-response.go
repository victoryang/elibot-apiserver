package apiv2

import (
	"net/http"
	//"fmt"
	"encoding/json"
)

type Response struct {
	Res 		string			`json:"res"`
}

func ToJson(res string) {
	response := Response{
		Res:	res,
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

func WriteInternalServerErrorResponse(w http.ResponseWriter, err error) {
	r := ToJson(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(r)
}