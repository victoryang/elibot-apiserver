package apiv1

import (
	"net/http"
	"os/exec"
	"io/ioutil"
	Log "elibot-apiserver/log"

	"github.com/gorilla/mux"
)

func rebootSystem(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("reboot")
	if err := cmd.Run(); err!=nil {
		Log.Error("reboot not executed: ", err)
		return
	}
	WriteSuccessResponse(w, "succeed in reboot")
}

func getSystemDate(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("date")
	stdout, err := cmd.StdoutPipe()
	cmd.Start()
	out, err := ioutil.ReadAll(stdout)
	if err != nil {
		Log.Error("output error: ", err)
		return
	}
	WriteJsonSuccessResponse(w, out)
}

func setSystemDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	args := " -s " + vars["date"]
	cmd := exec.Command("date", args)
	stdout, err := cmd.StdoutPipe()
	cmd.Start()
	out, err := ioutil.ReadAll(stdout)
	if err != nil {
		Log.Error("output error: ", err)
		return
	}
	WriteJsonSuccessResponse(w, out)
}