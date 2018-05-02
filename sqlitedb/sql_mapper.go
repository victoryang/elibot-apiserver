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
	M   *sqlitedb._Ctype_struct_sql_mapper
}

func Get_bookprogram_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))
    var sm *SqlMapper
	sm.m = C.get_bookprogram_sql_mapper(id)
	if sm == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(sm.m)
	return nil
}