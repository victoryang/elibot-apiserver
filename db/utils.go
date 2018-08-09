package db

import (
	"sync"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	Log "elibot-apiserver/log"
)

var DBName string
var mu sync.Mutex
var conn *sql.DB = nil

func createTableIfNotExist(table string) error {
	command := "CREATE TABLE IF NOT EXISTS " + table + "(key TEXT PRIMARY KEY NOT NULL, value TEXT NOT NULL)"
	if err := prepareAndExecuteCommand(command); err!=nil {
		Log.Error("create table fails: ", err)
	}
	return nil
}

func SetupDB(dbname string) {
	var err error
    DBName = dbname
    conn, err = sql.Open("sqlite3", dbname)
    if err!=nil {
    	Log.Error("failed to setup db")
    	return
    }
    if err := createTableIfNotExist(tableSettings); err!=nil {
    	Log.Error("create settings table fails")
    	return
    }
    return
}

func iterateRows(rows *sql.Rows) map[string]string {
	maps := make(map[string]string)
	for rows.Next() {
		var key string
		var value string
		if err := rows.Scan(&key, &value); err!=nil {
			Log.Error("scan rows fails: ", err)
			continue
		}
		maps[key] = value
	}
	return maps
}

func doQueryCommand(command string, args ...interface{}) (map[string]string, error){
	mu.Lock()
	defer mu.Unlock()
	rows, err := conn.Query(command, args...)
	if err!=nil {
		Log.Error("failed to query: ", err)
		return nil, err
	}
	return iterateRows(rows), nil
}

func prepareAndExecuteCommand(command string, args ...interface{}) error {
	mu.Lock()
	defer mu.Unlock()
	stmt, err := conn.Prepare(command)
	if err!=nil {
		Log.Error("failed to prepare statement")
		return err
	}
	_, err = stmt.Exec(args...)
	if err!=nil {
		Log.Error("failed to execute statement")
		return err
	}
	return nil
}