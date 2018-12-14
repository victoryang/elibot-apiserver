package apiv1

import (
	"net/http"
	"strconv"

	"elibot-apiserver/auth"
	Log "elibot-apiserver/log"
)

const (
	DefualtMode = 0
	RemoteModeId = "param.global.remoteAccessEnabled"
	cmdGetRemoteMode = "get_remote_mode_status"
)

var remoteMode int
var whitelist = make([]string, 0)

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

func InitAuthorization(wl []string) {
	whitelist = append(whitelist, wl...)
	InitRemoteModeFromParamServer()
}

func isRemoteReq(ip string) bool {
	for _, v := range whitelist {
		if ip == v {
			return true
		}
	}
	return false
}

func VerifyFunc(funcId string, authority int) bool {
	return true
}

func CHECKAUTHORIZATION(w http.ResponseWriter, r *http.Request, funcId string) bool {
	ip := getSourceIP(r)

	/* Reject remote request, if: 
	 * 1. remote mode not on
	 * 2. remote mode on, but no user login
	 */
	if isRemoteReq(ip) {
		if remoteModeOff() {
			WriteForbiddenResponse(w)
			return false
		} else {
			if !auth.CheckSession(ip) {
				WriteUnauthorizedResponse(w)
				return false
			}
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