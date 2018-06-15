package shm

// #cgo CFLAGS: -I/root/mcserver/include/
// #cgo LDFLAGS: -lshare
// #include <stdlib.h>
// #include <shared.h>
import "C"
import (
	"bytes"
	"sync"
	"encoding/json"
	"hash/crc32"

	Log "elibot-apiserver/log"
)

type SharedResource struct {
	Test 	int 		`json:"test,omitempty"`
}

var SharedResourcePool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, BufferSize))
	},
}
var crc_shared_resource int = 0

func getResourceAndCompare() (res []byte){
	mutex.Lock()
	defer mutex.Unlock()

	resource := SharedResource{
		Test:		1,
	}

	var crc int
	now, err := json.Marshal(resource)
	if err!=nil {
		crc = 0
	} else {
		crc = int(crc32.ChecksumIEEE(now))
	}

	var cache []byte
	buf := SharedResourcePool.Get().(*bytes.Buffer)
	if buf == nil {
		buf = bytes.NewBuffer(make([]byte, 0, BufferSize))
		crc_shared_resource = 0
	} else {
		cache = buf.Bytes()
	}

	if crc == 0 {
		crc_shared_resource = 0
 		res = []byte("")
 	} else {
 		if crc == crc_shared_resource {
 			res = nil
 			return
 		} else {
 			crc_shared_resource = crc
 			res = now
 		}
 	}

	buf.Reset()
	buf.Write(res)
	defer func(c []byte) {
		if e := recover(); e!=nil {
			Log.Error("buf write error: ", e)

			buf := make([]byte, 0, BufferSize*2)
			buf = c[:]
			SharedResourcePool = sync.Pool{
				New: func() interface{} {
					return bytes.NewBuffer(buf)
				},
			}
		}
	}(cache)

	SharedResourcePool.Put(buf)
	return
}

func watchSharedResource(modified chan []byte) {
	if res := getResourceAndCompare(); res != nil {
		modified <- res
	}
}

func InitSharedResource() int {
	return int(C.init_shared_resource())
}