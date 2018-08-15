package shm

// #cgo CFLAGS: -I../thirdparty/mrj/
// #cgo LDFLAGS: -L../thirdparty/mrj/ -lmrj
// #include <stdlib.h>
// #include <mrj.h>
import "C"
import (
	"sync"
	"hash/crc32"
	"unsafe"

	//Log "elibot-apiserver/log"
)

var crc_shared_resource int = 0
var resource_mutex sync.Mutex

func GetSysVar(datatype int, start int, end int) []byte {
	cstr := C.get_sysvar_data(C.int(datatype), C.int(start), C.int(end))
	gostr := C.GoString(cstr)
	defer C.free(unsafe.Pointer(cstr))
	return []byte(gostr)
}

func GetLocVar(datatype int, num int, start int, end int) []byte {
	cstr := C.get_locvar_data(C.int(datatype), C.int(num), C.int(start), C.int(end))
	gostr := C.GoString(cstr)
	defer C.free(unsafe.Pointer(cstr))
	return []byte(gostr)
}

func GetResourceOnce() []byte {
	cstr := C.get_resource_data()

	gostr := C.GoString(cstr)
	defer C.free(unsafe.Pointer(cstr))

	return []byte(gostr)
}

func getResourceAndCompare() []byte {
	mutex.Lock()
	defer mutex.Unlock()

	cstr := C.get_resource_data()

	gostr := C.GoString(cstr)
	defer C.free(unsafe.Pointer(cstr))

	buf := []byte(gostr)

	crc := int(crc32.ChecksumIEEE(buf))
	if crc == crc_shared_resource {
		//crc_shared_resource = 0
		return nil
	}
	crc_shared_resource = crc
	return buf
}

func watchSharedResource(modified chan []byte) {
	if res := getResourceAndCompare(); res != nil {
		modified <- res
	}
}

func InitSharedResource() int {
	return int(C.mrj_init())
}