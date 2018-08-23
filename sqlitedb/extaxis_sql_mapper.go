package sqlitedb

// #include<stdlib.h>
// #include<db/sqlmapper/extaxis_sql_mapper.h>
import "C"
import (
	"errors"

	Log "elibot-apiserver/log"
)

type ExtaxisSqlMapper struct {
	Id		string
}

func (m *ExtaxisSqlMapper) GetID() string {
	return m.Id
}

func (m *ExtaxisSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in ExtaxisSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_EXTAXIS_GET_ALL

	default:
		return errors.New("Not support")
	}

	return nil
}