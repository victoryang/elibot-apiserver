package auth

import (
	"elibot-apiserver/config"
)

func Init(c *config.Security) {
	SSLInit(c.SSLCert)
	SignatureInit(c.Sign)
}