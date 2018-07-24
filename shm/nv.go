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


var crc_nv int = 0
var mutex sync.Mutex

func getNVAndCompare() []byte {
	mutex.Lock()
	defer mutex.Unlock()

	cstr := C.get_nv_data()

	gostr := C.GoString(cstr)
	defer C.free(unsafe.Pointer(cstr))

	buf := []byte(gostr)

	crc := int(crc32.ChecksumIEEE(buf))
	if crc == crc_nv {
		//crc_nv = 0
		return nil
	}
	crc_nv = crc
	return buf
}

func watchNV(modified chan []byte) {
	if res := getNVAndCompare(); res != nil {
		modified <- res
	}
}

func InitNVRam() int {
	return int(C.init_nv_ram())
}