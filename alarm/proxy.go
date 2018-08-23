package alarm

import (
	"encoding/json"
)

type Response struct {
	Alarm		[]Record   		`json:"alarm"` 	
	TotalSize 	int 			`json:"TotalSize"`	
}

func GetRecords(start int, end int, timestamp uint32) ([]byte, error) {
	defer func() {
		clearUnReadRecordNumber()
	}()
	if end > len(records) {
		end = len(records)
	}

	ret := getRecordsByTimeStamp(timestamp, records[start:end])
	return json.Marshal(Response{Alarm: ret, TotalSize: len(ret)})
}

func GetRecordsByLevel(level string, start int, end int, timestamp uint32) ([]byte, error) {
	defer func() {
		clearUnReadRecordNumber()
	}()
	if end > len(records) {
		end = len(records)
	}

	tmp := getRecordsByTimeStamp(timestamp, records[start:end])
	var rec []Record
	for _, r := range tmp {
		if r.ErrNo[0] == level {
			rec = append(rec, r)
		}
	}
	return json.Marshal(Response{Alarm: rec, TotalSize: len(rec)}) 
}