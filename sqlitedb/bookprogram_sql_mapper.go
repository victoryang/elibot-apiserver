package sqlitedb

// #include<stdlib.h>
// #include<db/sqlmapper/bookprogram_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"

	Log "elibot-apiserver/log"
)

type BookProgramSqlMapper struct {
	Id         string
}

func (m *BookProgramSqlMapper) GetID() string {
	return m.Id
}

func (m *BookProgramSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in BookProgramSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_BOOKPROGRAM_GET_ALL

	default:
		return errors.New("Not support")
	}

	return nil
}