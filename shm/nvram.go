package shm

// #cgo CFLAGS: -I/root/mcserver/include/
// #cgo LDFLAGS: -lshare
// #include<stdlib.h>
// #include<sharedmemory.h>
import "C"
import (
	"fmt"
)

func checkNV() []byte {
	return []byte("checkNV")
}

func watchNV() string {
	value := C.watch_nv()
	return fmt.Sprint(uint64(value))
}