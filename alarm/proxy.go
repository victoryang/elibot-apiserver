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

	var recs []Record
	ret := getRecordsByTimeStamp(timestamp)
	length := len(ret)
	if length!=0 {
		if end > length {
			end = length 
		}
		recs = ret[start:end]
	} else {
		recs = nil
	}

	return json.Marshal(Response{Alarm: recs, TotalSize: length})
}

func GetRecordsByLevel(level string, start int, end int, timestamp uint32) ([]byte, error) {
	defer func() {
		clearUnReadRecordNumber()
	}()

	var recs []Record
	ret := getRecordsByTimeStamp(timestamp)

	var filter []Record
	for _, r := range ret {
		if r.ErrNo[0] == level {
			filter = append(filter, r)
		}
	}

	length := len(filter)
	if length!=0 {
		if end > length {
			end = length 
		}
		recs = filter[start:end]
	} else {
		recs = nil
	}
	return json.Marshal(Response{Alarm: recs, TotalSize: length}) 
}