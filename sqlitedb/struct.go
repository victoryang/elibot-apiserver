package sqlitedb

// #cgo LDFLAGS: -L/home/test/GitPro/sqlitedb -lsqlitedb
// #include<stdlib.h>
// #include<sqlmapper/bookprogram_sql_mapper.h>
import "C"

const (
	ELIBOT_BOOKPROGRAM_GET_ALL = "elibot.bookprogram.getAll"
)