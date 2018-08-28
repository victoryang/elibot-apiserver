package apiv1

import (
	"net/http"
	"io/ioutil"
	"encoding/json"

	Log "elibot-apiserver/log"
)

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