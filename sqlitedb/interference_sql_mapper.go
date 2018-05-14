package sqlitedb

// #include<stdlib.h>
// #include<include/interference_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"

	Log "elibot-apiserver/log"
)

type InterferenceSqlMapper struct {
	Id		string
}

func (m *InterferenceSqlMapper) register_interference_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	ism := C.get_interference_sql_mapper(id)
	if ism == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(ism)
	return nil
}

func (m *InterferenceSqlMapper) GetID() string {
	return m.Id
}

func (m *InterferenceSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Print("RegisterSqlMapper in InterferenceSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_INTERFERENCE_GET_ALL

	default:
		return errors.New("Not support")
	}

	return m.register_interference_sql_mapper(m.Id)
}