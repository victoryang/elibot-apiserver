package sqlitedb

// #include<stdlib.h>
// #include<include/zeropoint_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"

	Log "elibot-apiserver/log"
)

type ZeroPointSqlMapper struct {
	Id         string
}

func (m *ZeroPointSqlMapper) register_zeropoint_sql_mapper(q_id string) error {
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

func (m *ZeroPointSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Print("RegisterSqlMapper in ZeroPointSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_ZEROPOINT_GET_ALL

	default:
		return errors.New("Not support")
	}

	return m.register_zeropoint_sql_mapper(m.Id)
}