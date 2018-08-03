package sqlitedb

// #include<stdlib.h>
// #include<db/sqlmapper/io_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"

	Log "elibot-apiserver/log"
)

type IoSqlMapper struct {
	Id		string
}

func (m *IoSqlMapper) register_io_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	ism := C.get_io_sql_mapper(id)
	if ism == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(ism)
	return nil
}

func (m *IoSqlMapper) GetID() string {
	return m.Id
}

func (m *IoSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in IoSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_IO_GET_VALID_IOS_BY_GROUP

	default:
		return errors.New("Not support")
	}

	return m.register_io_sql_mapper(m.Id)
}