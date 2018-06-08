package sharedmemory

// #cgo CFLAGS: -I/root/mcserver/include/
// #cgo LDFLAGS: -lshare
// #include<stdlib.h>
// #include<sharedmemory.h>
import "C"
import(
	"fmt"
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
	res := string("Press reset value: ")+uint32(value)
	return []byte(res), true, nil
}

func watcher() {
	watchTicker := time.NewTicker(watchPeriod)
	defer func() {
		watchTicker.Stop()
	}()

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