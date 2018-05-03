package sqlitedb

// #include<stdlib.h>
// #include<sqlmapper/zeropoint_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"
)

type ZeroPointSqlMapper struct {
	Id         string
}

func (m *ZeroPointSqlMapper) get_zeropoint_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	bsm := C.get_zeropoint_sql_mapper(id)
	if bsm == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(bsm)
	return nil
}

func (m *ZeroPointSqlMapper) RegisterSqlMapper() error{
	m.Id = C.ELIBOT_ZEROPOINT_GET_ALL
	return m.get_zeropoint_sql_mapper(m.Id)
}