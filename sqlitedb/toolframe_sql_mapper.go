package sqlitedb

// #include<stdlib.h>
// #include<include/toolframe_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"

	Log "elibot-apiserver/log"
)

type ToolframeSqlMapper struct {
	Id		string
}

func (m *ToolframeSqlMapper) register_toolframe_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	tsm := C.get_common_toolframe_sql_mapper(id)
	if tsm == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(tsm)
	return nil
}

func (m *ToolframeSqlMapper) GetID() string {
	return m.Id
}

func (m *ToolframeSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in ToolframeSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_COMMON_GET_ALL_TOOLFRAMES

	default:
		return errors.New("Not support")
	}

	return m.register_toolframe_sql_mapper(m.Id)
}