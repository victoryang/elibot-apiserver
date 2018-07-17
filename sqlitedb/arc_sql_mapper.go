package sqlitedb

// #include <stdlib.h>
// #include <include/arc_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"

	Log "elibot-apiserver/log"
)

type ArcSqlMapper struct {
	Id		string
}

func (m *ArcSqlMapper) register_arc_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	asm := C.get_arc_sql_mapper(id)
	if asm == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(asm)
	return nil
}

func (m *ArcSqlMapper) GetID() string {
	return m.Id
}

func (m *ArcSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in ArcSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_ARC_GET_ALL_PARAMS

	case ELIBOT_ARC_GET_PARAMS:
		m.Id = C.ELIBOT_ARC_GET_PARAMS

	case ELIBOT_UPDATE_PARAMS:
		return errors.New("Not support now")
	default:
		return errors.New("Not support")
	}

	return m.register_arc_sql_mapper(m.Id)
}