package sqlitedb

// #include<stdlib.h>
// #include<sqlmapper/arc_sql_mapper.h>
import "C"
import (
	"fmt"
	"errors"
	"unsafe"
)

type ArcSqlMapper struct {
	Id		string
}

func (m *ArcSqlMapper) get_arc_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	asm := C.get_arc_sql_mapper(id)
	if asm == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(asm)
	return nil
}

func (m *ArcSqlMapper) GetID() string {
	return m.Id
}

func (m *ArcSqlMapper) RegisterSqlMapperForQueryAll() error{
	fmt.Println("RegisterSqlMapper in RegisterSqlMapperForQueryAll")
	m.Id = C.ELIBOT_ARC_GET_ALL_PARAMS
	return m.get_arc_sql_mapper(m.Id)
}

func (m *ArcSqlMapper) RegisterSqlMapper(mode int) error {
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		return m.RegisterSqlMapperForQueryAll()
	default:
		return errors.New("Not support")
	}
}