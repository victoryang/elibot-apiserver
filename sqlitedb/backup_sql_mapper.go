package sqlitedb

// #include<stdlib.h>
// #include<include/backup_sql_mapper.h>
import "C"
import (
	"errors"
	"unsafe"

	Log "elibot-apiserver/log"
)

type BackupSqlMapper struct {
	Id		string
}

func (m *BackupSqlMapper) get_backup_sql_mapper(q_id string) error {
	id := C.CString(q_id)
	defer C.free(unsafe.Pointer(id))

	ism := C.get_bak_sql_mapper(id)
	if ism == nil {
		return errors.New("Getting sqlmapper fails")
	}

	C.register_sql_mapper(ism)
	return nil
}

func (m *BackupSqlMapper) GetID() string {
	return m.Id
}

func (m *BackupSqlMapper) RegisterSqlMapperForQueryAll() error{
	m.Id = C.ELIBOT_BAK_BACKUP_PARAMS
	return m.get_backup_sql_mapper(m.Id)
}

func (m *BackupSqlMapper) RegisterSqlMapper(mode int) error {
	Log.Debug("RegisterSqlMapper in BackupSqlMapper | mode: ", mode)
	switch mode {
	case ELIBOT_GET_ALL_PARAMS:
		return m.RegisterSqlMapperForQueryAll()
	default:
		return errors.New("Not support")
	}
}