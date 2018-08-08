package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	Log "elibot-apiserver/log"
)

var DBName string
var conn *sql.DB = nil
var table = "elt_settings"

func SetupDB(dbname string) {
    DBName = dbname
    c, err := sql.Open("sqlite3", dbname)
    if err!=nil {
    	Log.Error("failed to setup db")
    	return
    }
    conn = c
}

func GetValue(key string) (string, error) {
	stmt, err := conn.Prepare("SELECT * FROM " + table + " where key=?")
	if err!=nil {
		Log.Error("failed to get value")
		return "", err
	}
	rows, err = stmt.Exec(key)
	var res string
	for rows.Next() {
		if err := rows.Scan(&res); err!=nil {
			continue
		}
	}
	return res, nil
}

func SetValue(key string, value string) error {
	stmt, err := conn.Prepare("update " + table + " set value=? where key=?")
	if err!=nil {
		Log.Error("failed to set value")
		return err
	}
	_, err = stmt.Exec(value, key)
	if err!=nil {
		Log.Error("failed to set value")
		return err
	}
	return nil
}