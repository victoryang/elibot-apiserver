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

type SqlMapper struct {
	m   *C.sql_mapper
}

func Get_bookprogram_sql_mapper(q_id string) (*SqlMapper, error) {
	id := C.CString(q_id)
    defer C.free(unsafe.Pointer(id))

    var mapper *SqlMapper
	mapper.m := C.get_bookprogram_sql_mapper(id)
	if mapper == nil {
		return nil, errors.New("Getting sqlmapper fails")
	}
	//sql.Mapper = mapper
	C.register_sql_mapper(mapper.m)
	return mapper, nil
}