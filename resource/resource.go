package resource

// #cgo CFLAGS: -I../build/include/include
// #cgo LDFLAGS: -L../build/lib -lrobresource
// #include <resource.h>
import "C"
import (
    "sync"
    "hash/crc32"
    "unsafe"
)

func GetNVOnce() []byte {
    cstr := C.elt_get_nv()

    gostr := C.GoString(cstr)
    defer C.free(unsafe.Pointer(cstr))

    return []byte(gostr)
}

func GetSysVar(datatype int, start int, end int) []byte {
    cstr := C.elt_get_sysvar(C.int(datatype), C.int(start), C.int(end))
    gostr := C.GoString(cstr)
    defer C.free(unsafe.Pointer(cstr))
    return []byte(gostr)
}

func GetLocVar(datatype int, num int, start int, end int) []byte {
    cstr := C.elt_get_locvar(C.int(datatype), C.int(num), C.int(start), C.int(end))
    gostr := C.GoString(cstr)
    defer C.free(unsafe.Pointer(cstr))
    return []byte(gostr)
}

func GetPLCOnce() []byte {
    cstr := C.elt_get_plc()

    gostr := C.GoString(cstr)
    defer C.free(unsafe.Pointer(cstr))

    return []byte(gostr)
}

func GetSharedOnce() []byte {
    cstr := C.elt_get_resource()

    gostr := C.GoString(cstr)
    defer C.free(unsafe.Pointer(cstr))

    return []byte(gostr)
}

var crc_nv int = 0
var crc_resource int = 0
var crc_plc int = 0

var nv_mutex sync.Mutex
var resource_mutex sync.Mutex
var plc_mutex sync.Mutex

func getNVAndCompare() []byte {
    nv_mutex.Lock()
    defer nv_mutex.Unlock()

    cstr := C.elt_get_nv()

    gostr := C.GoString(cstr)
    defer C.free(unsafe.Pointer(cstr))

    buf := []byte(gostr)

    crc := int(crc32.ChecksumIEEE(buf))
    if crc == crc_nv {
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

func getPLCAndCompare() []byte {
    plc_mutex.Lock()
    defer plc_mutex.Unlock()

    cstr := C.elt_get_plc()

    gostr := C.GoString(cstr)
    defer C.free(unsafe.Pointer(cstr))

    buf := []byte(gostr)

    crc := int(crc32.ChecksumIEEE(buf))
    if crc == crc_plc {
        return nil
    }
    crc_plc = crc
    return buf
}

func watchPLC(modified chan []byte) {
    if res := getPLCAndCompare(); res != nil {
        modified <- res
    }
}

func getResourceAndCompare() []byte {
    resource_mutex.Lock()
    defer resource_mutex.Unlock()

    cstr := C.elt_get_resource()

    gostr := C.GoString(cstr)
    defer C.free(unsafe.Pointer(cstr))

    buf := []byte(gostr)

    crc := int(crc32.ChecksumIEEE(buf))
    if crc == crc_resource {
        return nil
    }
    crc_resource = crc
    return buf
}

func watchResource(modified chan []byte) {
    if res := getResourceAndCompare(); res != nil {
        modified <- res
    }
}

func InitResource() int {
    return int(C.elt_resource_init())
}