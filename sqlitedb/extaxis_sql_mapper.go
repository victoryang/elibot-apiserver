package sqlitedb

// #include<stdlib.h>
// #include<sqlmapper/extaxis_sql_mapper.h>
import "C"
import (
	"fmt"
	"errors"
	"unsafe"
)

type ExtaxisSqlMapper struct {
	Id		string
}

func (m *ExtaxisSqlMapper) get_extaxis_sql_mapper(q_id string) error {
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

func (m *ExtaxisSqlMapper) RegisterSqlMapperForQueryAll() error{
	m.Id = C.ELIBOT_EXTAXIS_GET_ALL
	return m.get_extaxis_sql_mapper(m.Id)
}

func (m *ExtaxisSqlMapper) RegisterSqlMapper(mode int) error {
	fmt.Println("RegisterSqlMapper in ExtaxisSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		return m.RegisterSqlMapperForQueryAll()
	default:
		return errors.New("Not support")
	}
}