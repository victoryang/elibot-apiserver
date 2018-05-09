package sqlitedb

// #include<stdlib.h>
// #include<include/db/db_manager.h>
// #include<include/backup_sql_mapper.h>
import"C"
import (
	"unsafe"
	"errors"
)

const (
	DB_BACKUP = C.DB_BACKUP
	DB_RESTORE = C.DB_RESTORE
	DB_UPGRADE = C.DB_UPGRADE
)

type DBManager struct {
	mgr 		*C.db_manager
}

func (m *DBManager) Exec() error {
	m.mgr.commands = C.CString(C.ELIBOT_BAK_BACKUP_PARAMS)
	defer C.free(unsafe.Pointer(m.mgr.commands))

	op := make([]byte, 255)
	out := C.CBytes(op)
	defer C.free(out)

	res := C.DBMgrExecute(m.mgr, out)
	if res != C.DB_OK {
		return errors.New("DBManager: fail to execute the cmd")
	}
	return nil
}

func NewDBManager(DBname string, DBdir string, mode int) (*DBManager, error) {
    conn_str := C.CString(DBname)
	defer C.free(unsafe.Pointer(conn_str))

	db_dir := C.CString(DBdir)
	defer C.free(unsafe.Pointer(db_dir))

	db_mgr := new(DBManager)

    res := C.new_db_manager(
    		conn_str,
    		db_dir,
    		nil,
    		C.int(mode),
    		nil,
    		db_mgr.mgr,
    		C.char(0),
    	)
    if res != C.DB_OK {
    	return nil, errors.New("fail to new db manager")
    }
    return db_mgr, nil
}