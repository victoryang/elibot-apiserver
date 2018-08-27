package sqlitedb

// #include<stdlib.h>
// #include<db/sqlmapper/io_sql_mapper.h>
import "C"
import (
	"errors"

	Log "elibot-apiserver/log"
)

type IoSqlMapper struct {
	Id		string
}

func (m *IoSqlMapper) GetID() string {
	return m.Id
}

func (m *IoSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in IoSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_IO_GET_VALID_IOS_BY_GROUP

	default:
		return errors.New("Not support")
	}

	return nil
}