package settings

import (
	"elibot-apiserver/db"
	Log "elibot-apiserver/log"
)

var tableSettings = "elt_settings"

func GetAllKeyValue() (map[string]string, error) {
	command := "SELECT * FROM " + tableSettings
	return db.DoQueryCommand(command)
}

func GetKeyValue(key string) (map[string]string, error) {
	command := "SELECT * FROM " + tableSettings + " where key=?"
	return db.DoQueryCommand(command, key)
}

func SetKeyValue(key string, value string) error {
	command := "INSERT or REPLACE INTO " + tableSettings + " VALUES(?,?)"
	if err := db.PrepareAndExecuteCommand(command, key, value); err!=nil {
		Log.Error("failed to set value: ", err)
	}
	return nil
}

func SetupTable() {
	command := "CREATE TABLE IF NOT EXISTS " + tableSettings + "(key TEXT PRIMARY KEY NOT NULL, value TEXT NOT NULL)"
	if err := db.CreateTableIfNotExist(command); err!=nil {
    	Log.Error("create settings table fails")
    }
}