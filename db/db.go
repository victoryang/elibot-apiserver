package db

import (
	Log "elibot-apiserver/log"
)

var table = "elt_settings"

func GetValue(key string) (string, error) {
	return doQueryCommand("SELECT * FROM " + table + " where key=?", key)
}

func SetValue(key string, value string) error {
	command := "INSERT or REPLACE INTO " + table + " VALUES(?,?)"
	if err := prepareAndExecuteCommand(command, key, value); err!=nil {
		Log.Error("failed to set value: ", err)
	}
	return nil
}