package apiv1

import (
    "net/http"
    "elibot-apiserver/mcserver"
    "elibot-apiserver/paramserver"

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

func SendToParamServer(w http.ResponseWriter, serviceMethod string, params interface{}) {
    var reply string

    err := paramserver.SendToParamServerWithJsonRpc(serviceMethod, params, &reply)
    if err!=nil {
        Log.Error("Could not call rpc request to param server: ", err)
        WriteInternalServerErrorResponse(w, ERRMCSEVERNOTAVAILABLE)
        return
    }

    WriteJsonSuccessResponse(w, []byte(reply))
}

func InternalSendToParamServer(serviceMethod string, params interface{}, reply interface{}) error {
    if err := paramserver.SendToParamServerWithJsonRpc(serviceMethod, params, reply); err!=nil {
        Log.Error("Could not call internal request to param server: ", err)
        return err
    }

    return nil
}