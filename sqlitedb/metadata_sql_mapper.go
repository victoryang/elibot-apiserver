package sqlitedb

// #include<stdlib.h>
// #include<db/sqlmapper/metadata_sql_mapper.h>
import "C"
import (
	"errors"

	Log "elibot-apiserver/log"
)

type MetadataSqlMapper struct {
	Id		string
}

func (m *MetadataSqlMapper) GetID() string {
	return m.Id
}

func (m *MetadataSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in MetadataSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_METADATA_GET_ALL

	default:
		return errors.New("Not support")
	}

	return nil
}