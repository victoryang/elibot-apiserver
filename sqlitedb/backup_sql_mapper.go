package sqlitedb

// #include<stdlib.h>
// #include<db/sqlmapper/backup_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"

	Log "elibot-apiserver/log"
)

type BackupSqlMapper struct {
	Id		string
}

func (m *BackupSqlMapper) GetID() string {
	return m.Id
}

func (m *BackupSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in BackupSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		m.Id = C.ELIBOT_BAK_BACKUP_PARAMS
	default:
		return errors.New("Not support")
	}
	return nil
}