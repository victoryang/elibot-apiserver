package auth

import (
	"elibot-apiserver/config"
)

type JwtToken struct {
	PrivateKey 		string
	Signature 		string
}

var jwt *JwtToken = nil

func GetPrivateKey() string {
	return jwt.PrivateKey
}

func GetSignature() string {
	return jwt.Signature
}

func isJwt() bool {
	if jwt.PrivateKey == "" || jwt.Signature == "" {
		return false
	}
	return true
}

func JwtTokenInit(c *config.JwtToken) {
	jwt = new(JwtToken)

	jwt.PrivateKey = c.PrivateKey
	jwt.Signature = c.Signature
	return
}