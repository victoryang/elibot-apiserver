package sqlitedb

// #include<stdlib.h>
// #include<include/userframe_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"

	Log "elibot-apiserver/log"
)

type UserframeSqlMapper struct {
	Id		string
}

func (m *UserframeSqlMapper) register_userframe_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	usm := C.get_userframe_sql_mapper(id)
	if usm == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(usm)
	return nil
}

func (m *UserframeSqlMapper) GetID() string {
	return m.Id
}

func (m *UserframeSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in UserframeSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_USER_FRAME_GET_ALL

	default:
		return errors.New("Not support")
	}

	return m.register_userframe_sql_mapper(m.Id)
}