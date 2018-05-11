package middlewares

import (
	"github.com/urfave/negroni"
)

func AddAccesslog(n *Negroni) {
	alogger, err := accesslog.NewAccessLog(accesslogfile)
	if err!=nil {
		Log.Error("Error in opening access log file： ", err)
	} else {
		n.Use(alogger)
	}
}