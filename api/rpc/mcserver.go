package rpc

import (
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

func SendToMCServer(serviceMethod string, params interface{}, repch chan string) {
	var reply string

	err := mcserver.SendToMCServerWithJsonRpc(serviceMethod, params, &reply)
    if err!=nil {
        Log.Error("Could not call rpc request to mcserver: ", err)
        repch <- "Mcserver is not available right now"
        return
    }

    repch <- reply
}