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

type WsResponse struct {
	Alarm		[]Record   		`json:"alarm"` 	
	NewItemNo 	int 			`json:"NewItemNo"`	
}

func getRecordsByTimeStamp(time uint32, recs []Record) []Record {
	var rec []Record
	for _, r := range recs {
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
		if strings.Contains(r.Msg, "error") {
			ret = append(ret, r)
		}
	}

	rsp, err := json.Marshal(Response{Alarm: ret, NewItemNo: getUnReadRecordNumber()})
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
	clearUnReadRecordNumber()
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