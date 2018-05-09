package sqlitedb

// #include<stdlib.h>
// #include<include/extaxis_sql_mapper.h>
import "C"
import (
	"fmt"
	"errors"
	"unsafe"
)

type ExtaxisSqlMapper struct {
	Id		string
}

func (m *ExtaxisSqlMapper) register_extaxis_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	esm := C.get_extaxis_sql_mapper(id)
	if esm == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(esm)
	return nil
}

func (m *ExtaxisSqlMapper) GetID() string {
	return m.Id
}

func (m *ExtaxisSqlMapper) RegisterSqlMapper(mode int) error {
	fmt.Println("RegisterSqlMapper in ExtaxisSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_EXTAXIS_GET_ALL

	default:
		return errors.New("Not support")
	}

	return m.register_extaxis_sql_mapper(m.Id)
}