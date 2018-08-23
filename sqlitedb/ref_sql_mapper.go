package sqlitedb

// #include<stdlib.h>
// #include<db/sqlmapper/ref_sql_mapper.h>
import "C"
import (
	"errors"

	Log "elibot-apiserver/log"
)

type RefSqlMapper struct {
	Id		string
}

func (m *RefSqlMapper) GetID() string {
	return m.Id
}

func (m *RefSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in RefSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_REF_GET_ALL:
		m.Id = C.ELIBOT_REF_GET_ALL

	default:
		return errors.New("Not support")
	}

	return nil
}