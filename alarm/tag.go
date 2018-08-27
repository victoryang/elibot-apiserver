package alarm

import (
	"elibot-apiserver/db"
	Log "elibot-apiserver/log"
)

var tableSettings = "elt_settings"
var key = "alarm_tag"
var UnReadItemNo = 0

func getTagFromSettings() string {
	command := "SELECT * FROM " + tableSettings + " where key=?"
	if ret, err := db.DoQueryCommand(command, key); err==nil {
		return ret[key]
	}
	return ""
}

func setTagToSettings(value string) error {
	command := "INSERT or REPLACE INTO " + tableSettings + " VALUES(?,?)"
	if err := db.PrepareAndExecuteCommand(command, key, value); err!=nil {
		Log.Error("failed to set value: ", err)
	}

	return nil
}

func addNewRecordNumber(n int) {
	UnReadItemNo = UnReadItemNo + n
}

func getUnReadRecordNumber() int {
	return UnReadItemNo
}

func clearUnReadRecordNumber() {
	UnReadItemNo = 0
}