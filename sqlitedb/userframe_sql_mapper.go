package sqlitedb

// #include<stdlib.h>
// #include<db/sqlmapper/userframe_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"

	Log "elibot-apiserver/log"
)

type UserframeSqlMapper struct {
	Id		string
}

func (m *UserframeSqlMapper) GetID() string {
	return m.Id
}

func (m *UserframeSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in UserframeSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_USER_FRAME_GET_ALL

	default:
		return errors.New("Not support")
	}

	return nil
}