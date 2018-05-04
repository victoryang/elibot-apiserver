package sqlitedb

// #include<stdlib.h>
// #include<sqlmapper/userframe_sql_mapper.h>
import "C"
import (
	"fmt"
	"errors"
	"unsafe"
)

type UserframeSqlMapper struct {
	Id		string
}

func (m *UserframeSqlMapper) get_userframe_sql_mapper(q_id string) error {
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

func (m *UserframeSqlMapper) RegisterSqlMapper() error{
	fmt.Println("RegisterSqlMapper in UserframeSqlMapper")
	m.Id = C.ELIBOT_USER_FRAME_GET_ALL
	return m.get_userframe_sql_mapper(m.Id)
}

func (m *UserframeSqlMapper) RegisterSqlMapper(mode int) error {
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		return m.RegisterSqlMapperForQueryAll()
	default:
		return errors.New("Not support")
	}
}