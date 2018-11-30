package auth

import (
	"os"
	"elibot-apiserver/config"
)

var ssl *SSL = nil
var SSLEnable bool = false

type SSL struct {
	CaCert			string
	CertFile 		string
	KeyFile 		string
}

func GetCaCert() string {
	return ssl.CaCert
}

func GetCert() string {
	return ssl.CertFile
}

func GetKey() string {
	return ssl.KeyFile
}

func IsSSL() bool {
	return SSLEnable
}

func FileExists(path string) bool {
	st, e := os.Stat(path)
	// If file exists and is regular return true.
	if e == nil && st.Mode().IsRegular() {
		return true
	}
	return false
} 

func SSLInit(c *config.Certificate) {
	ssl = new(SSL)

	ssl.CaCert = c.Path + c.CaCert
	ssl.CertFile = c.Path + c.Certfile
	ssl.KeyFile = c.Path + c.Keyfile

	if FileExists(ssl.CertFile) && FileExists(ssl.KeyFile) && FileExists(ssl.CaCert) {
		SSLEnable =  true
	}

	return
}