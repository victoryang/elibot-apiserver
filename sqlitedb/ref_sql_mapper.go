package sqlitedb

// #include<stdlib.h>
// #include<sqlmapper/Ref_sql_mapper.h>
import "C"
import (
	"fmt"
	"errors"
	"unsafe"
)

type RefSqlMapper struct {
	Id		string
}

func (m *RefSqlMapper) get_ref_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	ism := C.get_ref_sql_mapper(id)
	if ism == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(ism)
	return nil
}

func (m *RefSqlMapper) GetID() string {
	return m.Id
}

func (m *RefSqlMapper) RegisterSqlMapperForQueryAll() error{
	m.Id = C.ELIBOT_REF_GET_ALL
	return m.get_ref_sql_mapper(m.Id)
}

func (m *RefSqlMapper) RegisterSqlMapper(mode int) error {
	fmt.Println("RegisterSqlMapper in RefSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		return m.RegisterSqlMapperForQueryAll()
	default:
		return errors.New("Not support")
	}
}