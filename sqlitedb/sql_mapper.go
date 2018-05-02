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

func Get_bookprogram_sql_mapper(q_id *C.char) (*SqlMapper, error) {
    var sm *SqlMapper
	sm.m = C.get_bookprogram_sql_mapper(id)
	if sm == nil {
		return nil, errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(sm.m)
	return sm, nil
}