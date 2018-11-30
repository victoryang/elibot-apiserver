package auth

import (
	"elibot-apiserver/config"
)

func AuthInit(c *config.Security) error {
	JwtTokenInit(c.Jwt)
	return InitManager()
}