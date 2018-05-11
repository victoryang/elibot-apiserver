package middleware

import (
	Log "elibot-apiserver/log"
	"elibot-apiserver/middleware/accesslog"
	"github.com/urfave/negroni"
)

func AddAccesslog(n *negroni.Negroni) {
	alogger, err := accesslog.NewLogger(accesslogfile)
	if err!=nil {
		Log.Error("Error in opening access log file： ", err)
	} else {
		n.Use(alogger)
	}
}