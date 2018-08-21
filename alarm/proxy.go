package alarm

import (
	"encoding/json"
)

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