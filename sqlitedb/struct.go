package sqlitedb

// #cgo LDFLAGS: -lsqlitedb
// #include<stdlib.h>
// #include<sqlmapper/bookprogram_sql_mapper.h>
import "C"

const (
	ELIBOT_BOOKPROGRAM_GET_ALL = C.ELIBOT_BOOKPROGRAM_GET_ALL
)