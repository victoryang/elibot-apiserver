package middleware

import (
	Log "elibot-apiserver/log"
	"elibot-apiserver/middleware/accesslog"
	"github.com/urfave/negroni"
)

func AddAccesslog(n *negroni.Negroni, accesslogfile string) *accesslog.Logger {
	logger, err := accesslog.NewLogger(accesslogfile)
	if err!=nil {
		Log.Error("Error in opening access log file： ", err)
		return nil
	} 
	return logger
}