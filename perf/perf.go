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
        AgentKey: "elibot-apiserver",
        AppName: "Basic Go Server",
        AppVersion: "1.0.0",
        AppEnvironment: "production",
    })

    http.HandleFunc(agent.ProfileHandlerFunc("/", handler))
    http.ListenAndServe(":9091", nil)
}