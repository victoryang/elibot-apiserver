package alarm

import (
	"encoding/json"
	"strings"
	"github.com/fsnotify/fsnotify"
	Log "elibot-apiserver/log"
	"elibot-apiserver/websocket"
)

var logfile = "/rbctrl/mcserver-err.log"
var ws *websocket.WsServer
var records []Record

type Response struct {
	Alarm		[]Record   		`json:"alarm"` 		
}

func GetRecordsByTimeStamp(time uint32) ([]byte, error) {
	var rec []Record
	for _, r := range records {
		if r.Time == time {
			rec = append(rec, r)
		}
	}
	return json.Marshal(Response{Alarm: rec}) 
}

func GetRecordsByErrNo(errno string) ([]byte, error) {
	var rec []Record
	for _, r := range records {
		if r.ErrNo[0] == errno {
			rec = append(rec, r)
		}
	}
	return json.Marshal(Response{Alarm: rec}) 
}

func GetRecordsNumber() int {
	return len(records)
}

func GetRecords(from int, end int) ([]byte, error) {
	if end > len(records) {
		return json.Marshal(Response{Alarm: records[from:]}) 
	}
	return json.Marshal(Response{Alarm: records[from:end]})
}

func GetAllRecords() ([]byte, error) {
	return json.Marshal(Response{Alarm: records})
}

func handleWriteEvent() {
	tmp := ScanAndParse(logfile)
	if len(tmp) == len(records) {
		return
	}

	rec := tmp[len(records):]
	records = tmp

	var ret []Record
	for _, r := range rec {
		if strings.Contains(r.Msg, "error") {
			ret = append(ret, r)
		}
	}

	rsp, err := json.Marshal(Response{Alarm: ret})
	if err!=nil {
		Log.Error("Could not marshal to json ", err)
	} else {
		ws.PushBytes(rsp)
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
	return
}

func NewAlarmMonitor(s *websocket.WsServer) error{
	InitRecords()
	fw, err := NewFileWatcher(logfile, AlarmHandler)
	if err!=nil {
		return err
	}
	fw.Watch()
	ws = s
	return nil
}