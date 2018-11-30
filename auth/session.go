package auth

import (
	//Log "elibot-apiserver/log"
)

type Session struct {
	Username 		string
	Token 			string
}

var loginSession = make(map[string]Session)

func SetSession(username string, ip string) string {
	token := createToken(username)

	session := Session {
		Username: 	username,
		Token:		token,
	}

	loginSession[ip] = session

	return token
}

func CheckSession(ip string) bool {
	_, ok := loginSession[ip]

	return ok
}

func VerifySessionToken(token string, ip string) bool {
	if _, ok := loginSession[ip]; !ok {
		return false
	}

	return loginSession[ip].Token == token
}

func ClearSession(ip string) {
	delete(loginSession, ip)
}