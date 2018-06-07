package sharedmemory

// #cgo CXXFLAGS: -I.
// #cgo LDFLAGS: -lshare
// #include<stdlib.h>
// #include<include/db/db_query.h>
import "C"
import(
	"byte"
	"fmt"
	"encoding/binary"
	"elibot-apiserver/api"
)

var wss *api.WsServer = nil

func getPressReset() {
	value := C.get_press_reset()
	fmt.Println("Get press reset value: ", value)
	buf := make([]byte, 4)
	binary.Write(buf, binary.LittleEndian, value.(int32))
	wss.Hub.broadcast <- buf
	return
}

func NewAndWatch(s *api.WsServer) {
	wss = s
}