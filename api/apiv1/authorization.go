package apiv1

import (
	"net/http"

	"elibot-apiserver/auth"
	"elibot-apiserver/db"
	Log "elibot-apiserver/log"
)

const (
	DefualtMode = 0
	RemoteModeId = "param.global.remoteAccessEnabled"
	cmdGetRemoteMode = "get_remote_mode_status"
)

var remoteMode int

func remoteModeOff() bool {
	return remoteMode == DefualtMode
}

func GetRemoteMode() int {
	return remoteMode
}

func SetRemoteMode(mode string) {
	v, e := strconv.Atoi(mode)
	if e == nil {
		remoteMode = v
	}
}

func InitRemoteModeFromParamServer() {
	var reply int32
	if err := InternalSendToParamServer(cmdGetRemoteMode, nil, &reply); err!=nil {
		Log.Error("init remote mode error, set to default")
		remoteMode = DefualtMode
		return
	}

	remoteMode = int(reply)
}

func InitAuthorization() error {
	InitRemoteModeFromParamServer()
}

func isRemoteReq(ip string) bool {
	return ip != "127.0.0.1"
}

func VerifyFunc(funcId string, authority id) bool {
	return true
}

func CHECKAUTHORIZATION(w http.ResponseWriter, r *http.Request, funcId string) bool {
	ip := getSourceIP(r)

	// Forbid remote ip, if remote mode not on
	if remoteModeOff() {
		if isRemoteReq(ip) {
			WriteForbiddenResponse(w)
			return false
		}
	}

	var authority int = 0
	if auth.CheckSession(ip) {
		token := getTokenFromHeader(r)
		if !auth.VerifySessionToken(token, ip) {
			WriteUnauthorizedResponse(w)
			return false
		}

		authority = auth.GetLoginedUserAuthority(ip)
	}

	if !VerifyFunc(funcId, authority) {
		WriteUnauthorizedResponse(w)
		return false
	}

	return true
}