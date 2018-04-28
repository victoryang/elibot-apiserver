package sqlitedb

// #cgo LDFLAGS: -lsqlitedb
// #include<stdlib.h>
// #include<db/db_context.h>

import "C"

const (
	ELIBOT_BOOKPROGRAM_GET_ALL = C.ELIBOT_BOOKPROGRAM_GET_ALL
)