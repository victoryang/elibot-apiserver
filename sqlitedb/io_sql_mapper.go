package sqlitedb

// #include<stdlib.h>
// #include<sqlmapper/io_sql_mapper.h>
import "C"
import (
	"fmt"
	"errors"
	"unsafe"
)

type IoSqlMapper struct {
	Id		string
}

func (m *IoSqlMapper) get_io_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	ism := C.get_io_sql_mapper(id)
	if ism == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(ism)
	return nil
}

func (m *IoSqlMapper) GetID() string {
	return m.Id
}

func (m *IoSqlMapper) RegisterSqlMapperForQueryAll() error{
	m.Id = C.ELIBOT_IO_GET_VALID_IOS_BY_GROUP
	return m.get_io_sql_mapper(m.Id)
}

func (m *IoSqlMapper) RegisterSqlMapper(mode int) error {
	fmt.Println("RegisterSqlMapper in IoSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		return m.RegisterSqlMapperForQueryAll()
	default:
		return errors.New("Not support")
	}
}