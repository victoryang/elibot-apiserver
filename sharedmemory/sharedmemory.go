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
var old []byte
const (
	duration = 5
	watchPeriod = time.Second * duration
)

func getPressReset() ([]byte, error) {
	value := C.get_press_reset()
	fmt.Println("Get press reset value: ", value)
	b := make([]byte, 2)
	buf := bytes.NewBuffer(b)
	err := binary.Write(buf, binary.LittleEndian, value)
	if err!=nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func watcher() {
	watchTicker := time.NewTicker(watchPeriod)
	for {
		select {
		case <-watchTicker.C:
			data, err := getPressReset()
			if err!=nil {
				fmt.Println(err)
			} else if data != old {
				old = data
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