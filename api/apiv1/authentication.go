package apiv1

import (
	"net/http"
	"strconv"

	"elibot-apiserver/auth"
	Log "elibot-apiserver/log"

	"github.com/gorilla/mux"
)

const (
	DefualtMode = 0

	AuthorityAdmin = 1
	AuthorityLogout = 5
	AuthorityRemoteExpert = 6
	AuthorityRemoteOperator = 7

	RemoteModeId = "param.global.remoteAccessEnabled"
	cmdGetRemoteMode = "get_remote_mode_status"
)

var remoteMode int
var whitelist = make([]string, 0)

var redirectFuncIdMap = map[string]string{
	"param.softlimit.limit_min_pos": "PARAM0200030",
	"param.softlimit.limit_max_pos": "PARAM0200030",
	"param.speed.speed_max_joint": "PARAM0200040",
	"param.speed.speed_min_joint": "PARAM0200040",
	"param.global.remoteAccessEnabled": "PARAM0200060",
	"param.system.lcd_backlight": "SYSSETTING130020010",
	"param.system.lcd_time": "SYSSETTING130020020",
}

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

func isRemoteAuthority(authority int) bool {
	if authority == AuthorityRemoteExpert || authority == AuthorityRemoteOperator {
		return true
	}

	return false
}

func InitRemoteModeFromParamServer() {
	var reply int32
	if err := InternalSendToParamServer(cmdGetRemoteMode, nil, &reply); err!=nil {
		Log.Error("init remote mode error: ", err)
		remoteMode = DefualtMode
		return
	}

	Log.Debug("init remote mode: ", reply)
	remoteMode = int(reply)
}

func InitAuthorization(wl []string) {
	whitelist = append(whitelist, wl...)
	InitRemoteModeFromParamServer()
}

func isRemoteReq(ip string) bool {
	for _, v := range whitelist {
		if ip == v {
			return false
		}
	}
	return true
}

func VerifyFunc(funcId string, authority int) bool {
	return true
}

func MustVerify(funcId string) bool {
	if funcId == "" {
		return false
	}

	return true
}

func MustRedirectFuncId(funcId string) bool {
	ret := false
	switch funcId {
		case "setParam":
			ret = true
		default:
	}

	return ret
}

func redirectFuncId(funcId *string, r *http.Request) {
	id := mux.Vars(r)["md_id"]
	if redirect, ok := redirectFuncIdMap[id]; ok {
		*funcId = redirect
	}
}

func authenticator(w http.ResponseWriter, r *http.Request, next http.Handler) {
	funcId := mux.CurrentRoute(r).GetName()

	if funcId == "login" || funcId == "logout" {
		next.ServeHTTP(w, r)
		return
	}

	ip := getSourceIP(r)

	var authority int = AuthorityLogout
	if token := getTokenFromHeader(r); token!="" {
		if !auth.CheckSession(token) {
			WriteUnauthorizedResponse(w)
			return
		}

		authority = auth.GetLoginedUserAuthority(token)
	}

	/* Reject remote request, if:
	 * 1. remote mode not on
	 * 2. remote mode on, but not a logined user
	 */
	if isRemoteReq(ip) {
		if remoteModeOff() {
			WriteForbiddenResponse(w)
			return
		} else {
			if authority == AuthorityLogout {
				WriteUnauthorizedResponse(w)
				return
			}
		}
	}

	if MustVerify(funcId) {
		Log.Debug("currentRoute is: ", funcId)

		if MustRedirectFuncId(funcId) {
			redirectFuncId(&funcId, r)
			Log.Debug("redirect to route ", funcId)
		}

		if !VerifyFunc(funcId, authority) {
			WriteUnauthorizedResponse(w)
			return
		}
	}

	next.ServeHTTP(w, r)
}

type AuthenticationMiddleware struct {
}

func NewAuthenticationMiddleware() *AuthenticationMiddleware {
	return &AuthenticationMiddleware{}
}

func (awm *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authenticator(w, r, next)
	})
}