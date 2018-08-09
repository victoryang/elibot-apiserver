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

func SetupDB(dbname string) {
	var err error
    DBName = dbname
    if conn, err = sql.Open("sqlite3", dbname); err!=nil {
    	Log.Error("failed to setup db: ", err)
    }
    return
}

func CreateTableIfNotExist(table string) error {
	command := "CREATE TABLE IF NOT EXISTS " + table + "(key TEXT PRIMARY KEY NOT NULL, value TEXT NOT NULL)"
	if err := PrepareAndExecuteCommand(command); err!=nil {
		Log.Error("create table fails: ", err)
	}
	return nil
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

func DoQueryCommand(command string, args ...interface{}) (map[string]string, error){
	mu.Lock()
	defer mu.Unlock()
	rows, err := conn.Query(command, args...)
	if err!=nil {
		Log.Error("failed to query: ", err)
		return nil, err
	}
	return iterateRows(rows), nil
}

func PrepareAndExecuteCommand(command string, args ...interface{}) error {
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