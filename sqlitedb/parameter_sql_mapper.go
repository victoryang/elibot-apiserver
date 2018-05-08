package sqlitedb

// #include<stdlib.h>
// #include<sqlmapper/parameter_sql_mapper.h>
import "C"
import (
	"fmt"
	"errors"
	"unsafe"
)

type ParameterSqlMapper struct {
	Id		string
}

func (m *ParameterSqlMapper) get_parameter_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	ism := C.get_params_sql_mapper(id)
	if ism == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(ism)
	return nil
}

func (m *ParameterSqlMapper) GetID() string {
	return m.Id
}

func (m *ParameterSqlMapper) RegisterSqlMapperForQueryWithParams() error {
	m.Id = C.ELIBOT_PARAMS_GET_VALID_PARAM_BY_ID
	return m.get_parameter_sql_mapper(m.Id)
}

func (m *ParameterSqlMapper) RegisterSqlMapperForQueryAll() error{
	m.Id = C.ELIBOT_ENUM_GET_ALL
	return m.get_parameter_sql_mapper(m.Id)
}

func (m *ParameterSqlMapper) RegisterSqlMapper(mode int) error {
	fmt.Println("RegisterSqlMapper in ParameterSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		return m.RegisterSqlMapperForQueryAll()
	case ELIBOT_GET_WITH_PARAMS:
		return m.RegisterSqlMapperForQueryWithParams()
	default:
		return errors.New("Not support")
	}
}