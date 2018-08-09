package db

import (
	Log "elibot-apiserver/log"
)

var tableSettings = "elt_settings"

func GetAllSettingsKV() (map[string]string, error) {
	command := "SELECT * FROM " + tableSettings
	return doQueryCommand(command)
}

func GetSettingsValue(key string) (map[string]string, error) {
	command := "SELECT * FROM " + tableSettings + " where key=?"
	return doQueryCommand(command, key)
}

func SetSettingsValue(key string, value string) error {
	command := "INSERT or REPLACE INTO " + tableSettings + " VALUES(?,?)"
	if err := prepareAndExecuteCommand(command, key, value); err!=nil {
		Log.Error("failed to set value: ", err)
	}
	return nil
}