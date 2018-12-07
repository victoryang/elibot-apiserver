package auth

import (
	"elibot-apiserver/config"
	Log "elibot-apiserver/log"

	"github.com/dgrijalva/jwt-go"
)

var Signature string

func createToken(username string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims.(jwt.MapClaims)["username"] = username

	tokenString, err := token.SignedString([]byte(Signature))
	if err != nil {
		Log.Error("Failed to sign token: ", err)
	}

	return tokenString
}

func getSignature() string {
	return Signature
}

func JwtTokenInit(c *config.JwtToken) {
	Signature = c.Signature
}