package sharedmemory

// #cgo CFLAGS: -I/root/mcserver/include/
// #cgo LDFLAGS: -lshare
// #include<stdlib.h>
// #include<sharedmemory.h>
import "C"
import(
	"bytes"
	"fmt"
	"encoding/binary"
	"elibot-apiserver/api"
)

var wss *api.WsServer = nil

func getPressReset() error {
	value := C.get_press_reset()
	fmt.Println("Get press reset value: ", value)
	b := make([]byte, 1)
	buf := bytes.NewBuffer(b)
	binary.Write(buf, binary.LittleEndian, value)
	if hub, ok := wss.Hub.(PushMsg); ok {
		hub.PushMsg(buf.Bytes())
	}
	return nil
}

func NewAndWatch(s *api.WsServer) {
	wss = s
	return
}