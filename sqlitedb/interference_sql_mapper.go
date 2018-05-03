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
	BaseSqlMapper
}

func (m *InterferenceSqlMapper) get_interference_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	bsm := C.get_interference_sql_mapper(id)
	if bsm == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(bsm)
	return nil
}

func (m *InterferenceSqlMapper) RegisterSqlMapper() error{
	fmt.Println("RegisterSqlMapper in InterferenceSqlMapper")
	m.Id = C.ELIBOT_INTERFERENCE_GET_ALL
	return m.get_interference_sql_mapper(m.Id)
}