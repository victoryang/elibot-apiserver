package apiv1

import (
	"net/http"

	"elibot-apiserver/auth"
	//Log "elibot-apiserver/log"
)

type AuthorityTree map[string]int

var remoteMode bool

func GetRemoteMode() bool {
	return remoteMode
}

func SetRemoteMode(mode bool) {
	remoteMode = mode
}

func InitAuthorization() {
	// initialize remote mode value from db

	// initialize authorization tree from mcserver
}

func isLocalReq(ip string) bool {
	return ip == "127.0.0.1"
}

func checkAuthorization(authority string, funcName string) bool {
	return true
}

func CHECKAUTHORIZATION(r *http.Request, funcName string) bool {
	ip := getSourceIP(r)

	// Forbid remote ip, if remote mode not on
	if !GetRemoteMode() {
		if !isLocalReq(ip) {
			return false
		}
	}

	var authority int = 0
	if auth.CheckSession(ip) {
		token := getTokenFromHeader(r)
		if !auth.VerifySessionToken(token, ip) {
			return false
		}

		authority = auth.GetLoginedUserAuthority(ip)
	}

	// TODO
	if !checkAuthorization(authority, funcName) {
		return false
	}

	return true
}