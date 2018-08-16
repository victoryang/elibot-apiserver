package fs

import (
	"os/exec"
	"io/ioutil"
	"encoding/json"
	"strings"
	"github.com/fsnotify/fsnotify"
	Log "api-server/log"
	"api-server/websocket"
)

var logfile = "/rbctrl/mcserver-err.log"
var ws websocket.WsServer

type Response struct {
	serverlog 		string 		`json:"serverlog,omitempty"`
}

func readLastLineFromFile() string {
	cmd := exec.Command("tail", "-n 1", logfile)
	stdout, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err!=nil {
		Log.Error("Failed to exec cmd: ", err)
		return ""
	}
	content, err := ioutil.ReadAll(stdout)
	if err!=nil {
		Log.Error("Error to read stdout ", err)
		return ""
	}
	return string(content)
}

func handleWriteEvent() {
	res := readLastLineFromFile()
	if strings.Contains(res, "error") {
		rsp,_ := json.Marshal(Response{
			serverlog:	res,
			})
		ws.PushBytes(rsp)
	}
}

func logWatcherHandler(evt fsnotify.Event, err error) {
	if err!=nil {
		Log.Debug("watch error: ", err)
		return
	}

	Log.Debug("server log changed: ", evt.Op)
	switch evt.Op {
	case fsnotify.Write:
		handleWriteEvent()
	case fsnotify.Remove:
	}
}

func NewMCServerLogWatcher(s websocket.WsServer) error{
	fw, err := NewFileWatcher(logfile, logWatcherHandler)
	if err!=nil {
		return err
	}
	fw.Watch()
	ws = s
	return nil
}