package sqlitedb

// #include<stdlib.h>
// #include<sqlmapper/toolframe_sql_mapper.h>
import "C"
import (
	"fmt"
	"errors"
	"unsafe"
)

type ToolframeSqlMapper struct {
	Id		string
}

func (m *ToolframeSqlMapper) get_toolframe_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	tsm := C.get_toolframe_sql_mapper(id)
	if tsm == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(tsm)
	return nil
}

func (m *ToolframeSqlMapper) GetID() string {
	return m.Id
}

func (m *ToolframeSqlMapper) RegisterSqlMapper() error{
	fmt.Println("RegisterSqlMapper in ToolframeSqlMapper")
	m.Id = C.ELIBOT_COMMON_GET_ALL_TOOLFRAMES
	return m.get_toolframe_sql_mapper(m.Id)
}

func (m *ToolframeSqlMapper) RegisterSqlMapper(mode int) error {
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		return m.RegisterSqlMapperForQueryAll()
	default:
		return errors.New("Not support")
	}
}