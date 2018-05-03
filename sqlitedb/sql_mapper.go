package sqlitedb

// #cgo LDFLAGS: -lsqlitedb
// #include<stdlib.h>
// #include<sql_mapper.h>
// #include<sqlmapper/bookprogram_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"
)

type SqlMapper interface {
	RegisterSqlMapper(id string) ()
}

type BaseSqlMapper struct {
}