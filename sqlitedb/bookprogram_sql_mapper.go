package sqlitedb

// #include<stdlib.h>
// #include<sqlmapper/bookprogram_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"
)

type BookProgramSqlMapper struct {
	Id         string
}

func (m *BookProgramSqlMapper) get_bookprogram_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	bsm := C.get_bookprogram_sql_mapper(id)
	if bsm == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(bsm)
	return nil
}

func (m *BookProgramSqlMapper) RegisterSqlMapperForQueryAll() error{
	fmt.Println("RegisterSqlMapperForQueryAll in BookProgramSqlMapper")
	m.Id = C.ELIBOT_BOOKPROGRAM_GET_ALL
	return m.get_bookprogram_sql_mapper(m.Id)
}

func (m *BookProgramSqlMapper) RegisterSqlMapper(mode int) error {
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		return m.RegisterSqlMapperForQueryAll()
	default:
		return errors.New("Not support")
	}
}