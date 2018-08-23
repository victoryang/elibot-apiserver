package sqlitedb

// #include<stdlib.h>
// #include<db/sqlmapper/zeropoint_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"

	Log "elibot-apiserver/log"
)

type ZeroPointSqlMapper struct {
	Id         string
}

func (m *ZeroPointSqlMapper) GetID() string {
	return m.Id
}

func (m *ZeroPointSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in ZeroPointSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_ZEROPOINT_GET_ALL

	default:
		return errors.New("Not support")
	}

	return nil
}