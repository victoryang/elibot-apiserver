package auth

import (
	"fmt"
	"elibot-apiserver/config"
	Log "elibot-apiserver/log"

	"github.com/dgrijalva/jwt-go"
)

var Signature string

func createToken(username string, ip string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims.(jwt.MapClaims)["username"] = username
	token.Claims.(jwt.MapClaims)["ip"] = ip

	tokenString, err := token.SignedString([]byte(Signature))
	if err != nil {
		Log.Error("Failed to sign token: ", err)
	}

	return tokenString
}

func getClaimsFromToken(tokenString string, username *string, ip *string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	        return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	    }

	     return []byte(Signature), nil
	})

	if err!=nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		*username = claims["username"].(string)
		*ip = claims["ip"].(string)
		return nil
	}

	return fmt.Errorf("token valid: ", token.Valid)
}

func JwtTokenInit(c *config.JwtToken) {
	Signature = c.Signature
}