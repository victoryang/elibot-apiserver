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
)

type SharedResource struct {
	Test 	int 		`json:"test,omitempty"`
}

var SharedResourcePool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, bufferSize))
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
	}
	crc = int(crc32.ChecksumIEEE(now))
	res = now

	buf := SharedResourcePool.Get().(*bytes.Buffer)
	if buf == nil {
		if crc == 0 {
			res = []byte("")
			return
		} else {
			crc_shared_resource = 0
		}
	}

	cache := buf.Bytes()
	if crc == crc_shared_resource {
		SharedResourcePool.Put(buf)
		res = nil
		return
	}

	buf.Reset()
	crc_shared_resource = crc
	buf.Write(now)
	defer func(c []byte) {
		if e := recover(); e!=nil {
			Log.Error("buf write error: ", e)
			crc_shared_resource = 0
			SharedResourcePool = sync.Pool{
				New: func() interface{} {
					return bytes.NewBuffer(make([]byte, 0, bufferSize*2))
				},
			}
			buf.Write(c)
			res = c
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

func InitSharedResource() {
	C.init_shared_resource()
}