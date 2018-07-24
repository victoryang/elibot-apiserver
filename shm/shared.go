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