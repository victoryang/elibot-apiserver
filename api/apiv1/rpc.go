package apiv1

import (
    "net/http"
    "elibot-apiserver/mcserver"

    Log "elibot-apiserver/log"
)

func ConcatParams(args ...string) []string {
	var params []string
	for _,v := range args {
		if v == "" {
			continue
		}
		params = append(params, v)
	}

	return params
}

func SendToMCServer(w http.ResponseWriter, serviceMethod string, params interface{}) {
	var reply string

	err := mcserver.SendToMCServerWithJsonRpc(serviceMethod, params, &reply)
    if err!=nil {
        Log.Error("Could not call rpc request to mcserver: ", err)
        WriteInternalServerErrorResponse(w, ERRMCSEVERNOTAVAILABLE)
        return
    }

	WriteJsonSuccessResponse(w, []byte(reply))
}

func InternalSendToMCServer(serviceMethod string, params interface{}, reply interface{}) error {
    if err := mcserver.SendToMCServerWithJsonRpc(serviceMethod, params, reply); err!=nil {
        Log.Error("Could not call internal request to mcserver: ", err)
        return err
    }

    return nil
}