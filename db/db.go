package db

import (
	Log "elibot-apiserver/log"
)

var table = "elt_settings"

func GetValue(key string) (map[string]string, error) {
	command := "SELECT * FROM " + table + " where key=?"
	return doQueryCommand(command, key)
}

func SetValue(key string, value string) error {
	command := "INSERT or REPLACE INTO " + table + " VALUES(?,?)"
	if err := prepareAndExecuteCommand(command, key, value); err!=nil {
		Log.Error("failed to set value: ", err)
	}
	return nil
}