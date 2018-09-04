package perf

import (
	"net/http"
	"fmt"
	stackimpact "github.com/stackimpact/stackimpact-go"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world!")
}

func StartPerfMonitor() {
	agent := stackimpact.Start(stackimpact.Options{
        AgentKey: "e32976f8baa8ee4d669221dc2e39d2264f18156c",
        AppName: "MyGoApp",
        AppVersion: "1.0.0",
        AppEnvironment: "production",
        DashboardAddress:":5050",
    })

    http.HandleFunc(agent.ProfileHandlerFunc("/stackimpact/", handler))
    http.ListenAndServe(":9091", nil)
}