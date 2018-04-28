package db

import (
	"fmt"
	"errors"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const SQLITE3 = "sqlite3"

type DB struct {
	db     *sql.DB
}

func (d *DB)CheckAvailandAccess() (err error){
	err = d.db.Ping()
	return
}

func (d *DB)OpenDB(name string) error {
	db, err := sql.Open(SQLITE3, name)
    if err != nil {
        fmt.Println("open fail")
        return err
    }
    d.db = db
    return nil
}

func (d *DB)CloseDB() {
	d.Close()
}