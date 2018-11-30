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

func AddUserWithTransaction(name string, mail string, mobile string, authority int, secretkey string) bool {
	mu.Lock()
	defer mu.Unlock()

	if name == "" || secretkey == "" {
		return false
	}

	tx, err := conn.Begin()
	if err!=nil {
		Log.Error("failed to begin transaction")
		return false
	}
	defer tx.Rollback()

	stmtU, err1 := tx.Prepare(userAddCommand)
	if err1!=nil {
		Log.Error("failed to prepare stmtU: ", err1)
		return false
	}
	defer stmtU.Close()

	if _, err2 := stmtU.Exec(name, mail, mobile, authority); err2!=nil {
		Log.Error("failed to exec stmtU: ", err2)
		return false
	}

	stmtS, err3 := tx.Prepare(secretAddCommand)
	if err!=nil {
		Log.Error("failed to prepare stmtS: ", err3)
		return false
	}
	defer stmtS.Close()

	if _, err4 := stmtS.Exec(name, secretkey); err4!=nil {
		Log.Error("failed to exec stmtS: ", err4)
		return false
	}

	if err = tx.Commit(); err!=nil {
		return false
	}

	return true
}

func RemoveUserWithTransaction(name string) bool {
	mu.Lock()
	defer mu.Unlock()

	if name == "" {
		return false
	}

	tx, err := conn.Begin()
	if err!=nil {
		Log.Error("failed to begin transaction")
		return false
	}
	defer tx.Rollback()

	stmtS, err1 := tx.Prepare(secretRemoveCommand)
	if err1!=nil {
		Log.Error("failed to prepare stmtS: ", err1)
		return false
	}
	defer stmtS.Close()

	if _, err2 := stmtS.Exec(name); err2!=nil {
		Log.Error("failed to exec stmtS: ", err2)
		return false
	}

	stmtU, err3 := tx.Prepare(userRemoveCommand)
	if err!=nil {
		Log.Error("failed to prepare stmtU: ", err3)
		return false
	}
	defer stmtU.Close()

	if _, err4 := stmtU.Exec(name); err4!=nil {
		Log.Error("failed to exec stmtU: ", err4)
		return false
	}

	if err = tx.Commit(); err!=nil {
		return false
	}

	return true
}