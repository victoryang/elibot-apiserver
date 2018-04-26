package db

import (
	"fmt"
	"errors"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DBContext struct {
	dbname    string
	pool      DBConnPool
}

func DBContextCancel() {

}

func NewDBContext(dbname string) (*DBContext, error) {
	ctx := new(DBContext)
	ctx.dbname = dbname
	
	db := new(DB)
	db.OpenDB(ctx.dbname)
	db.CloseDB()

	return ctx, nil
}