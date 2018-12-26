package auth

import (
	//Log "elibot-apiserver/log"
)

type Session struct {
	Username 		string
	Authority 		int
	IP 				string
}

// loginSession   [token]Session
var loginSession = make(map[string]Session)

func GetLoginedUserAuthority(token string) int {
	return loginSession[token].Authority
}

func GetLoginedUserInfo(token string) Session {
	return loginSession[token]
}

func SetSession(username string, authority int, ip string) string {
	token := createToken(username, ip)
	session := Session {
		Username: 	username,
		Authority: 	authority,
		IP:			ip,
	}

	loginSession[token] = session

	return token
}

func CheckSession(token string) bool {
	_, ok := loginSession[token]

	return ok
}

func ClearSession(token string) {
	delete(loginSession, token)
}