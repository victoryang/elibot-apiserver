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
	conn        *C.char
	dir         *C.char
}

func (m *DBManager) Exec() error {
	m.mgr.commands = C.CString(C.ELIBOT_BAK_BACKUP_PARAMS)
	defer C.free(unsafe.Pointer(m.mgr.commands))

	op := make([]byte, 255)
	out := C.CBytes(op)
	defer C.free(out)

	res := C.DBMgrExecute(m.mgr, out)
	C.free(unsafe.Pointer(m.conn))
	C.free(unsafe.Pointer(m.dir))
	if res != C.DB_OK {
		return errors.New("DBManager: fail to execute the cmd")
	}
	return nil
}

func NewDBManager(DBname string, DBdir string, mode int) (*DBManager, error) {
	mgr := new(DBManager)
	mgr.mgr = C.NewDBManager()

	mgr.conn = C.CString(DBname)
	mgr.dir = C.CString(DBdir)

    res := C.new_db_manager(
    		mgr.conn,
    		mgr.dir,
    		nil,
    		C.int(mode),
    		nil,
    		mgr.mgr,
    		C.char(0),
    	)
    if res != C.DB_OK {
    	return nil, errors.New("fail to new db manager")
    }
    return mgr, nil
}