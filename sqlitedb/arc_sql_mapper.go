package sqlitedb

// #include <stdlib.h>
// #include <db/sqlmapper/arc_sql_mapper.h>
import "C"
import (
	"errors"

	Log "elibot-apiserver/log"
)

type ArcSqlMapper struct {
	Id			string
}

func (m *ArcSqlMapper) GetID() string {
	return m.Id
}

func (m *ArcSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in ArcSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_ARC_GET_ALL_PARAMS

	case ELIBOT_ARC_GET_PARAMS:
		m.Id = C.ELIBOT_ARC_GET_PARAMS

	case ELIBOT_UPDATE_PARAMS:
		return errors.New("Not support now")
	default:
		return errors.New("Not support")
	}

	return nil
}