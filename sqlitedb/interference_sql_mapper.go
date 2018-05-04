package sqlitedb

// #include<stdlib.h>
// #include<sqlmapper/interference_sql_mapper.h>
import "C"
import (
	"fmt"
	"errors"
	"unsafe"
)

type InterferenceSqlMapper struct {
	Id		string
}

func (m *InterferenceSqlMapper) get_interference_sql_mapper(q_id string) error {
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

func (m *InterferenceSqlMapper) RegisterSqlMapperForQueryAll() error{
	fmt.Println("RegisterSqlMapper in RegisterSqlMapperForQueryAll")
	m.Id = C.ELIBOT_INTERFERENCE_GET_ALL
	return m.get_interference_sql_mapper(m.Id)
}

func (m *InterferenceSqlMapper) RegisterSqlMapper(mode int) error {
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		return m.RegisterSqlMapperForQueryAll()
	default:
		return errors.New("Not support")
	}
}