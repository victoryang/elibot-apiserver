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

	zsm := C.get_zeropoint_sql_mapper(id)
	if zsm == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(zsm)
	return nil
}

func (m *ZeroPointSqlMapper) GetID() string {
	return m.Id
}

func (m *ZeroPointSqlMapper) RegisterSqlMapperForQueryAll() error{
	m.Id = C.ELIBOT_ZEROPOINT_GET_ALL
	return m.get_zeropoint_sql_mapper(m.Id)
}

func (m *ZeroPointSqlMapper) RegisterSqlMapper(mode int) error {
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		return m.RegisterSqlMapperForQueryAll()
	default:
		return errors.New("Not support")
	}
}