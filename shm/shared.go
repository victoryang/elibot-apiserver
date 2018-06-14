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
		return new(bytes.Buffer)
	},
}
var crc_shared_resource int = 0

func getResourceAndCompare() []byte{
	resource := SharedResource{
		Test:		1,
	}
	buf := SharedResourcePool.Get().(*bytes.Buffer)
	old := buf.Bytes()
	now, err := json.Marshal(resource)
	if err!=nil {
		if old!=nil {
			return old
		} else {
			return nil
		}
	}
	crc_now := int(crc32.ChecksumIEEE(now))
	if crc_now != crc_shared_resource {
		buf.Reset()
		_, err := buf.Write(now)
		if err != nil {
			crc_shared_resource = 0
			return nil
		}
		buf.Reset()
		SharedResourcePool.Put(buf.Bytes())
		crc_shared_resource = crc_now
		return now
	}
	return nil
}

func watchSharedResource(modified chan []byte) {
	modified <- getResourceAndCompare()
}

func InitSharedResource() {
	C.init_shared_resource()
}