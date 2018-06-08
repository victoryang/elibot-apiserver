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
	"time"
	"elibot-apiserver/api"
)

var wss *api.WsServer = nil
var old _Ctype_int
const (
	duration = 5
	watchPeriod = time.Second * duration
)

func getPressResetIfModified() ([]byte, bool, error) {
	value := C.get_press_reset()
	fmt.Println("get value ", value)
	if old == value {
		return nil, false, nil
	}
	old = value
	b := make([]byte, 2)
	buf := bytes.NewBuffer(b)
	err := binary.Write(buf, binary.LittleEndian, value)
	return buf.Bytes(), true, err
}

func watcher() {
	watchTicker := time.NewTicker(watchPeriod)
	for {
		select {
		case <-watchTicker.C:
			data, modified, err := getPressResetIfModified()
			if err!=nil {
				fmt.Println(err)
			} else if modified {
				wss.Hub.PushMsg(data)
			}
		}
	}
}

func handleMsg(msg []byte) {
	fmt.Println(string(msg[:]))
}

func NewAndWatch(s *api.WsServer) {
	wss = s
	wss.Hub.NotificationRegister(handleMsg)
	go watcher()
	return
}