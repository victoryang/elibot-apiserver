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

func getAndCompare() []byte{
	resource := SharedResource{
		ProjectName: 		getMainFile(),
		CurCoordinate:		getCurCoordinate(),
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
		SharedResourcePool.Put(buf.Write(now))
		crc_shared_resource = crc_now
		return now
	}
	return nil
}

func watchSharedResource(modified chan []byte) {
	modified <- getAndCompare()
}

func InitSharedResource() {
	C.init_shared_resource()
}