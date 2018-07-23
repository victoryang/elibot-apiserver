package sqlitedb

// #include<stdlib.h>
// #include<mcmanager.h>
import "C"
import (
	"unsafe"
	Log "elibot-apiserver/log"
)

func SqlitedbBackup(DBname string, DBdir string) int {
	conn := C.CString(DBname)
	defer C.free(unsafe.Pointer(conn))

	dir := C.CString(DBdir)
	defer C.free(unsafe.Pointer(dir))

	res := C.mcmanager_backup_db(conn, dir)
	if res != C.ERRNONE {
		Log.Error("fail to backup db")
    	return int(res)
    }
    return 0
}

func SqlitedbRestore(DBname string, DBdir string, backupname string, force_restore int) int {
	conn := C.CString(DBname)
	defer C.free(unsafe.Pointer(conn))

	dir := C.CString(DBdir)
	defer C.free(unsafe.Pointer(dir))

	bakname := C.CString(backupname)
	defer C.free(unsafe.Pointer(bakname))

	res := C.mcmanager_restore_db(conn, dir, bakname, C.char(force_restore))
	if res != C.ERRNONE {
		Log.Error("fail to restore db")
    	return int(res)
    }
    return 0
}