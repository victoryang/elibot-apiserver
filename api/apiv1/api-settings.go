package apiv1

import (
	"net"
	"net/http"
	"os/exec"
	"io/ioutil"
	Log "elibot-apiserver/log"

	"github.com/gorilla/mux"
)

func runShell(cmd *exec.Cmd, w http.ResponseWriter) {
	stdout, _ := cmd.StdoutPipe()
    err := cmd.Start()
	if err!=nil {
		Log.Error("Failed to exec cmd: ", err)
		WriteInternalServerErrorResponse(w, ERRRUNCMD)
		return
	}
	content, err := ioutil.ReadAll(stdout)
	if err!=nil {
		Log.Error("Error to read stdout ", err)
		WriteInternalServerErrorResponse(w, ERRRUNCMD)
		return
	}
	WriteJsonSuccessResponse(w, content)
}

func rebootSystem(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("reboot")
	err := cmd.Start()
	if err!=nil {
		Log.Error("Failed to reboot: ", err)
		WriteInternalServerErrorResponse(w, ERRRUNCMD)
		return
	}
	WriteSuccessResponse(w, "succeed in reboot")
}

func getSystemDate(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("date")
	runShell(cmd, w)
}

func setSystemDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cmd := exec.Command("date", "-s", vars["date"])
	runShell(cmd, w)
}

func getSystemIP(w http.ResponseWriter, r *http.Request) {
	addrs, err := net.InterfaceAddrs()
	if err!=nil {
		Log.Error("Failed to get any address: ", err)
		WriteInternalServerErrorResponse(w, ERRRUNCMD)
		return
	}

	var LocalIP string
	for _, address := range addrs {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                LocalIP = ipnet.IP.To4()
            }
        }
    }

    WriteSuccessResponse(w, LocalIP)
}