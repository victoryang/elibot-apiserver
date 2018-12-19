package alarm

import (
	"encoding/json"
	"strings"
	"github.com/fsnotify/fsnotify"
	Log "elibot-apiserver/log"
	"elibot-apiserver/websocket"
)

var logfile = "/rbctrl/mcserver-err.log"
var serviceMethod = "push_message_to_network"
var records []Record

type WsResponse struct {
	Alarm		[]Record   		`json:"alarm"` 	
	NewItemNo 	int 			`json:"NewItemNo"`	
}

func getRecordsByTimeStamp(time uint32) []Record {
	var rec []Record
	for _, r := range records {
		if r.Time >= time {
			rec = append(rec, r)
		}
	}
	return rec 
}

func handleWriteEvent() {
	tmp := ScanAndParse(logfile)
	if len(tmp) == len(records) {
		return
	}

	rec := tmp[len(records):]
	addNewRecordNumber(len(rec))

	records = tmp

	var ret []Record
	for _, r := range rec {
		if strings.Contains(r.ErrNo[0], "0") {
			ret = append(ret, r)
		}
	}

	message, err := json.Marshal(WsResponse{Alarm: ret, NewItemNo: getUnReadRecordNumber()})
	if err!=nil {
		Log.Error("Could not marshal to json ", err)
	} else {
		var reply bool
		params := make(map[string]interface{})
		params["message"] = string(message)

	    err := paramserver.SendToParamServerWithJsonRpc(serviceMethod, params, &reply)
	    if err!=nil {
	        Log.Error("Could not call rpc request to param server: ", err)
	        return
	    }

	    Log.Debug("reply is ", reply)
	}
	return
}

func AlarmHandler(evt fsnotify.Event, err error) {
	if err!=nil {
		Log.Debug("watch error: ", err)
		return
	}

	Log.Debug("server log changed: ", evt.Op)
	switch evt.Op {
	case fsnotify.Write:
		handleWriteEvent()
	}
}

func InitRecords() {
	records = ScanAndParse(logfile)
	clearUnReadRecordNumber()
	return
}

func NewAlarmMonitor() error {
	InitRecords()
	fw, err := NewFileWatcher(logfile, AlarmHandler)
	if err!=nil {
		return err
	}
	fw.Watch()
	return nil
}