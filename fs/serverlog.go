package fs

import (
	"os"
	"encoding/json"
	"strings"
	"github.com/fsnotify/fsnotify"
	Log "elibot-apiserver/log"
	"elibot-apiserver/websocket"
)

var logfile = "/rbctrl/mcserver-err.log"
var ws *websocket.WsServer

type Response struct {
	serverlog 		string 		`json:"serverlog,omitempty"`
}

func readLastLineFromFile() string {
	file, err := os.Open(logfile)
	if err!=nil {
		return ""
	}

	r, err := newReverseReader(file)
	if err!= nil {
		return ""
	}

	p := make([]byte, 127)
	n, err := r.Read(p)
	if err!= nil {
		return ""
	}

	res := strings.Split(string(p[:n]), "\n")
	if res == nil {
		return ""
	}

	for i:=len(res)-1; i>=0; i-- {
		if res[i] != "" {
			return res[i]
		}
	}

	return ""
}

func handleWriteEvent() {
	res := readLastLineFromFile()
	if res == "" {
		return
	}
	if strings.Contains(res, "error") {
		rsp,err := json.Marshal(Response{serverlog:	res})
		if err!=nil {
			Log.Error("Could not marshal to json ", err)
		} else {
			ws.PushBytes(rsp)
		}
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

func NewMCServerLogWatcher(s *websocket.WsServer) error{
	fw, err := NewFileWatcher(logfile, logWatcherHandler)
	if err!=nil {
		return err
	}
	fw.Watch()
	ws = s
	return nil
}