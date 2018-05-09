package sqlitedb

// #include<stdlib.h>
// #include<include/sqlmapper/parameter_sql_mapper.h>
import "C"
import (
	"fmt"
	"errors"
	"unsafe"
)

type ParameterSqlMapper struct {
	Id		string
}

func (m *ParameterSqlMapper) register_parameter_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	psm := C.get_params_sql_mapper(id)
	if psm == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(psm)
	return nil
}

func (m *ParameterSqlMapper) GetID() string {
	return m.Id
}

func (m *ParameterSqlMapper) RegisterSqlMapper(mode int) error {
	fmt.Println("RegisterSqlMapper in ParameterSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_ENUM_GET_ALL

	case ELIBOT_PARAMETER_GET_PARAMS:
		m.Id = C.ELIBOT_PARAMS_GET_PARAMS

	case ELIBOT_PARAMETER_GET_BY_ID:
		m.Id = C.ELIBOT_PARAMS_GET_VALID_PARAM_BY_ID

	case ELIBOT_PARAMETER_GET_BY_GROUP:
		m.Id = C.ELIBOT_PARAMS_GET_VALID_PARAMS_BY_GROUP

	case ELIBOT_REF_GET_ALL:

	default:
		return errors.New("Not support")
	}

	return m.register_parameter_sql_mapper(m.Id)
}