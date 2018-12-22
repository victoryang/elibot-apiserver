package apiv1

import (
	"net/http"
	Log "elibot-apiserver/log"

	"github.com/gorilla/mux"
)

const (
	cmdGenerateMachineCode = "generate_machine_code"
	cmdGetMachineCode = "get_machine_code"
	cmdEncrypt = "encrypt_system"
	cmdDecrypt = "input_lisence"
)

func getEncryptionStatus(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get encryption status")
	var status int
	SendToParamServerWithReply(w, "get_encryption_status", nil, &status)
}

func getEncryptionRemainTime(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get encryption remain time")
	var remaintime int
	SendToParamServerWithReply(w, "get_encryption_remain_time", nil, &remaintime)
}

func generateMachineCode(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting generate machine code")

	SendToMCServer(w, cmdGenerateMachineCode, nil)
}

func getMachineCode(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting get machine code")

	SendToMCServer(w, cmdGetMachineCode, nil)
}

func doEncrypt(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting encrypt")

	vars := mux.Vars(r)
	id := vars["id"]
	SendToMCServer(w, cmdEncrypt, ConcatParams(id))
}

func doDecrypt(w http.ResponseWriter, r *http.Request) {
	Log.Debug("starting decrypt")

	vars := mux.Vars(r)
	lisence := vars["lisence"]
	SendToMCServer(w, cmdDecrypt, ConcatParams(lisence))
}
