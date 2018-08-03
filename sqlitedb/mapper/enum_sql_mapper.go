package sqlitedb

// #include<stdlib.h>
// #include<db/sqlmapper/enum_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"

	Log "elibot-apiserver/log"
)

type EnumSqlMapper struct {
	Id		string
}

func (m *EnumSqlMapper) register_enum_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	ism := C.get_enum_sql_mapper(id)
	if ism == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(ism)
	return nil
}

func (m *EnumSqlMapper) GetID() string {
	return m.Id
}

func (m *EnumSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in EnumSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_ENUM_GET_ALL

	default:
		return errors.New("Not support")
	}

	return m.register_enum_sql_mapper(m.Id)
}