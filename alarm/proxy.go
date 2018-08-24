package alarm

import (
	"encoding/json"
	Log "elibot-apiserver/log"
)

type Response struct {
	Alarm		[]Record   		`json:"alarm"` 	
	TotalSize 	int 			`json:"TotalSize"`	
}

func getReverseRecords(input []Record, start int, end int, length int) []Record {
	if length == 0 {
		return nil
	}

	if end > length {
		end = length 
	}

	var recs []Record
	s := length-1-start
	e := length-1-end

	for i:=s; i>e;i-- {
		recs = append(recs, input[i])
	}
	return recs
}

func GetRecords(start int, end int, timestamp uint32) ([]byte, error) {
	defer func() {
		clearUnReadRecordNumber()
	}()

	ret := getRecordsByTimeStamp(timestamp)
	length := len(ret)

	if start > length {
		Log.Error("index start out of range")
		return json.Marshal(Response{Alarm: nil, TotalSize: length})
	}

	recs := getReverseRecords(ret, start, end, length)

	return json.Marshal(Response{Alarm: recs, TotalSize: length})
}

func GetRecordsByLevel(level string, start int, end int, timestamp uint32) ([]byte, error) {
	defer func() {
		clearUnReadRecordNumber()
	}()

	ret := getRecordsByTimeStamp(timestamp)

	var filter []Record
	for _, r := range ret {
		if r.ErrNo[0] == level {
			filter = append(filter, r)
		}
	}

	length := len(filter)
	if start > length {
		Log.Error("index start out of range")
		return json.Marshal(Response{Alarm: nil, TotalSize: length})
	}

	recs := getReverseRecords(filter, start, end, length)
	return json.Marshal(Response{Alarm: recs, TotalSize: length}) 
}