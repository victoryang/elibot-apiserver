package sqlitedb

// #include<stdlib.h>
// #include<sqlmapper/ref_sql_mapper.h>
import "C"
import (
	"fmt"
	"errors"
	"unsafe"
)

type RefSqlMapper struct {
	Id		string
}

func (m *RefSqlMapper) register_ref_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	rsm := C.get_ref_sql_mapper(id)
	if rsm == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(rsm)
	return nil
}

func (m *RefSqlMapper) GetID() string {
	return m.Id
}

func (m *RefSqlMapper) RegisterSqlMapper(mode int) error {
	fmt.Println("RegisterSqlMapper in RefSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_REF_GET_ALL:
		m.Id = C.ELIBOT_REF_GET_ALL

	default:
		return errors.New("Not support")
	}

	return m.register_ref_sql_mapper(m.Id)
}