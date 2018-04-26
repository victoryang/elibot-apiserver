package db

import(
	"context"
	"errors"
	"fmt"
	"database/sql"
)

const DB_POOL_SIZE_DEFAULT = 1

type DBConnPool struct {
	db       *DB
	conn     []*sql.Conn
	size     int
	ctx      context
}

func CreateDBConnPool(s string, size int) (pool *DBConnPool, err error){
	pool = new(DBConnPool)
	pool.db = new(DB)
	err = pool.db.OpenDB(s)
	if err!=nil {
		pool=nil
		return
	}
	defer pool.db.CloseDB()
	pool.size = size
	err = nil
	return
}