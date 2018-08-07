package apiv1

import (
	"net/http"
	"os/exec"
	Log "elibot-apiserver/log"

	"github.com/gorilla/mux"
)

func rebootSystem(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("reboot")
	if err := cmd.Run(); err!=nil {
		Log.Error("reboot not executed: ", err)
	}
	WriteSuccessResponse(w, "succeed in reboot")
}

func getSystemDate(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("date")
	out, err := cmd.Output()
	if err != nil {
		Log.Error(err)
	}
	WriteJsonSuccessResponse(w, out)
}

func setSystemDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]
	command := "date -s " + date
	cmd := exec.Command(command)
	out, err := cmd.Output()
	if err != nil {
		Log.Error(err)
	}
	WriteJsonSuccessResponse(w, out)
}