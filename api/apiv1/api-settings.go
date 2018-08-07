package apiv1

import (
	"net/http"
	"os/exec"
	Log "elibot-apiserver/log"

	//"github.com/gorilla/mux"
)

func rebootSystem(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("reboot")
	if err := cmd.Run(); err!=nil {
		Log.Error("reboot not executed")
	}
	WriteSuccessResponse(w, "succeed in reboot")
}