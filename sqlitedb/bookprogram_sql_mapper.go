package sqlitedb

// #include<stdlib.h>
// #include<sql_mapper.h>
// #include<sqlmapper/bookprogram_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"
)

type BookProgramSqlMapper struct {
	Id         string
}

func (m *BookProgramSqlMapper) get_bookprogram_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	bsm := C.get_bookprogram_sql_mapper(id)
	if bsm == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(bsm)
	return nil
}

func (m *BookProgramSqlMapper) RegisterSqlMapper() error{
	m.Id = ELIBOT_BOOKPROGRAM_GET_ALL
	return m.get_bookprogram_sql_mapper(m.Id)
}