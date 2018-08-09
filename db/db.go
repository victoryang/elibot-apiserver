package db

import (
	Log "elibot-apiserver/log"
)

var table = "elt_settings"

func GetValue(key string) (string, error) {
	return doQueryCommand("SELECT * FROM " + table + " where key=?", key)
}

func SetValue(key string, value string) error {
	if err := prepareAndExecuteCommand("INSERT or REPLACE INTO " + table + " VALUES(?,?)", key, value); err!=nil {
		Log.Error("failed to set value: ", err)
	}
	return nil
}