package sqlitedb

// #include<stdlib.h>
// #include<include/metadata_sql_mapper.h>
import "C"
import (
	"fmt"
	"errors"
	"unsafe"
)

type MetadataSqlMapper struct {
	Id		string
}

func (m *MetadataSqlMapper) register_metadata_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	ism := C.get_metadata_sql_mapper(id)
	if ism == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(ism)
	return nil
}

func (m *MetadataSqlMapper) GetID() string {
	return m.Id
}

func (m *MetadataSqlMapper) RegisterSqlMapper(mode int) error {
	fmt.Println("RegisterSqlMapper in MetadataSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_METADATA_GET_ALL

	default:
		return errors.New("Not support")
	}

	return m.register_metadata_sql_mapper(m.Id)
}