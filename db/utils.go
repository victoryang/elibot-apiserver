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

func createTableIfNotExist() error {
	if err := prepareAndExecuteCommand("CREATE TABLE IF NOT EXISTS elt_settings(key TEXT PRIMARY KEY NOT NULL, value TEXT NOT NULL)"); err!=nil {
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
    if err := createTableIfNotExist(); err!=nil {
    	Log.Error("create table fails: ", err)
    	return
    }
    return
}

func iterateRows(rows *sql.Rows) string {
	var res []string
	var r string
	for rows.Next() {
		if err := rows.Scan(&r); err!=nil {
			continue
		}
		res = append(res, r)
	}
	if res == nil {
		return ""
	}
	return res[0]
}

func doQueryCommand(command string, args ...interface{}) (string, error){
	mu.Lock()
	defer mu.Unlock()
	rows, err := conn.Query(command, args)
	if err!=nil {
		Log.Error("failed to query: ", err)
		return "", err
	}
	return iterateRows(rows), nil
}

func prepareAndExecuteCommand(command string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	stmt, err := conn.Prepare(command)
	if err!=nil {
		Log.Error("failed to prepare statement")
		return err
	}
	_, err = stmt.Exec(args)
	if err!=nil {
		Log.Error("failed to execute statement")
		return err
	}
	return nil
}