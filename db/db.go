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
    if conn, err = sql.Open("sqlite3", dbname+"?_journal_mode=WAL"); err!=nil {
    	Log.Error("failed to setup db: ", err)
    }
    return
}

func CloseDB() error {
	return conn.Close()
}

func InitializeTable(table string) error {
	if err := PrepareAndExecuteCommand("CREATE TABLE IF NOT EXISTS " + table); err!=nil {
		Log.Error("Initialize table fails: ", err)
	}
	return nil
}

func CreateTableIfNotExist(command string) error {
	if err := PrepareAndExecuteCommand(command); err!=nil {
		Log.Error("create table fails: ", err)
	}
	return nil
}

func DoQuery(command string, args ...interface{}) (*sql.Rows, error){
	mu.Lock()
	defer mu.Unlock()
	return conn.Query(command, args...)
}

func PrepareAndExecuteCommand(command string, args ...interface{}) error {
	mu.Lock()
	defer mu.Unlock()
	stmt, err := conn.Prepare(command)
	if err!=nil {
		Log.Error("failed to prepare statement")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err!=nil {
		Log.Error("failed to execute statement")
		return err
	}
	return nil
}