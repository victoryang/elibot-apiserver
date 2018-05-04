package sqlitedb

// #include<stdlib.h>
// #include<sqlmapper/enum_sql_mapper.h>
import "C"
import (
	"fmt"
	"errors"
	"unsafe"
)

type EnumSqlMapper struct {
	Id		string
}

func (m *EnumSqlMapper) get_enum_sql_mapper(q_id string) error {
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

func (m *EnumSqlMapper) RegisterSqlMapperForQueryAll() error{
	m.Id = C.ELIBOT_ENUM_GET_ALL
	return m.get_enum_sql_mapper(m.Id)
}

func (m *EnumSqlMapper) RegisterSqlMapper(mode int) error {
	fmt.Println("RegisterSqlMapper in EnumSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		return m.RegisterSqlMapperForQueryAll()
	default:
		return errors.New("Not support")
	}
}