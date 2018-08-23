package sqlitedb

// #include<stdlib.h>
// #include<db/sqlmapper/toolframe_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"

	Log "elibot-apiserver/log"
)

type ToolframeSqlMapper struct {
	Id		string
}

func (m *ToolframeSqlMapper) GetID() string {
	return m.Id
}

func (m *ToolframeSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in ToolframeSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_COMMON_GET_ALL_TOOLFRAMES

	default:
		return errors.New("Not support")
	}

	return nil
}