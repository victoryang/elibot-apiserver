package main

import (
	"elibot-apiserver/config"
)

type Signature struct {
	AccessKey 		string
	SecretKey 		string
}

var sign *Signature

func SignatureInit(c *config.Signature) {
	sign = new(Signature)

	sign.AccessKey = c.AccessKey
	sign.SecretKey = c.SecretKey
	return
}