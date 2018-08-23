package sqlitedb

// #include<stdlib.h>
// #include<db/sqlmapper/interference_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"

	Log "elibot-apiserver/log"
)

type InterferenceSqlMapper struct {
	Id		string
}

func (m *InterferenceSqlMapper) GetID() string {
	return m.Id
}

func (m *InterferenceSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in InterferenceSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_INTERFERENCE_GET_ALL

	default:
		return errors.New("Not support")
	}

	return nil
}