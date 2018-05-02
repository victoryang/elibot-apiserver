package sqlitedb

// #cgo LDFLAGS: -lsqlitedb
// #include<stdlib.h>
// #include<sql_mapper.h>
// #include<sqlmapper/bookprogram_sql_mapper.h>
import "C"
import (
	"errors"
)

type SqlMapper struct {
	m   *C.sql_mapper
}

func Get_bookprogram_sql_mapper(q_id string) (*SqlMapper, error) {
	id := C.CString(q_id)
    defer C.free(unsafe.Pointer(id))

	mapper := C.get_bookprogram_sql_mapper(id)
	if mapper == 0 {
		return nil, errors.Error("Getting sqlmapper fails")
	}
	//sql.Mapper = mapper
	C.register_sql_mapper(mapper.m)
	return mapper, nil
}