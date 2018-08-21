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

func GetAlreadyReadTag() int {
	if ret, err := strconv.Atoi(getTagFromSettings()); err==nil {
		return ret
	}
	return 0
}

func GetRecordsNumber() int {
	return len(records)
}

func GetRecords(from int, end int) ([]byte, error) {
	if end > len(records) {
		end = len(records)
	}
	return json.Marshal(Response{Alarm: records[from:end]})
}

func GetAllRecords() ([]byte, error) {
	return json.Marshal(Response{Alarm: records})
}