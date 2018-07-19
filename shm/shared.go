package shm

// #cgo CFLAGS: -I../thirdparty/mrj/
// #cgo LDFLAGS: -L../thirdparty/mrj/ -lmrj
// #include <stdlib.h>
// #include <mrj.h>
import "C"
import (
	"sync"
	"encoding/json"
	"hash/crc32"
	"unsafe"

	Log "elibot-apiserver/log"
)

type SharedResource struct {
	Resource 			string 			`json:"resource,omitempty"`
}

var crc_shared_resource int = 0
var resource_mutex sync.Mutex

func getResourceAndCompare() (res []byte){
	mutex.Lock()
	defer mutex.Unlock()

	cj := C.get_resource_data()

	cstr := C.cJSON_PrintUnformatted(cj)
    defer C.cJSON_Delete(cj)

    gostr := C.GoString(cstr)
    defer C.free(unsafe.Pointer(cstr))

	resource := SharedResource{
		Resource:	reformatString(gostr),
	}

	buf, err := json.Marshal(resource)
	if err!=nil {
		Log.Error("Json marshal error: ", err)
		crc_shared_resource = 0
		return []byte("Json marshal error: " + err.Error())
	}

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