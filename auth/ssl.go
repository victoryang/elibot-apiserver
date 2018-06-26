package auth

import (
	"crypto/tls"
	"elibot-apiserver/config"
)

var ssl *SSL
var SSLEnable bool = false

type SSL struct {
	CertFile 		string
	KeyFile 		string
}

func GetCert() {
	return ssl.CertFile
}

func GetKey() {
	return ssl.KeyFile
}

func IsSSL() bool {
	return SSLEnable
}

func SSLInit(c *config.Certificate) {
	ssl = new(SSL)

	ssl.CertFile = c.Path + c.Certfile
	ssl.KeyFile = c.Path + c.Keyfile

	if FileExists(ssl.CertFile) && FileExists(ssl.KeyFile) {
		SSLEnable =  true
	}

	return
}