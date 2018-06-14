package shm

// #cgo CFLAGS: -I/root/mcserver/include/
// #cgo LDFLAGS: -lshare
// #include<stdlib.h>
// #include "include/shared.h"
import "C"
import (
	"fmt"
	"sync"
	"encoding/json"
)

type SharedResource struct {
	
}

var NVRamPool = sync.Pool{
	New: func() interface{} {
		return new(SharedResource)
	}
}

func getAxisCount() int32 {
	return int32(C.get_axis_count());
}

func getMainFile() string {
	return C.GoString(C.get_main_file())
}

func getCurCoordinate() int32{
	return int32(C.get_cur_coordinate())
}

func getNV() []byte{
	nvram := NvRam{
		ProjectName: 		getMainFile(),
		CurCoordinate:		getCurCoordinate(),
	}
	b, _ := json.Marshal(nvram)
	return b
}

func watchSharedResource() ([]byte, string) {
	value := C.watch_nv()
	return getNV(), fmt.Sprint(uint64(value))
}

func InitSharedResource() {
	C.init_shared_resource()
}