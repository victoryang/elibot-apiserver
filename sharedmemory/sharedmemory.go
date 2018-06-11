package sharedmemory

// #cgo CFLAGS: -I/root/mcserver/include/
// #cgo LDFLAGS: -lshare
// #include<stdlib.h>
// #include<sharedmemory.h>
import "C"
import(
	"context"
	"fmt"
	"time"
	"elibot-apiserver/api"
)

var wss *api.WsServer = nil
var ctx_shm context.Context
var cancel context.CancelFunc
var old _Ctype_int
const (
	duration = 5
	watchPeriod = time.Second * duration
)

func getPressResetIfModified() ([]byte, bool, error) {
	value := C.get_press_reset()
	fmt.Println("get value ", value)
	/*if old == value {
		return nil, false, nil
	}*/
	old = value
	res := fmt.Sprint(uint64(value))
	return []byte(res), true, nil
}

func getZeroEncodeIfModified() ([]byte, bool, error) {
	value := C.get_zero_encode(0)
	fmt.Println("get zero encode: ", value)
	/*if old == value {
		return nil, false, nil
	}*/
	old = value
	res := fmt.Sprint(uint64(value))
	return []byte(res), true, nil
}

func watcher(ctx context.Context) {
	watchTicker := time.NewTicker(watchPeriod)
	defer func() {
		watchTicker.Stop()
	}()

	for {
		select {
		case <-watchTicker.C:
			data, modified, err := getZeroEncodeIfModified()
			if err!=nil {
				fmt.Println(err)
			} else if modified {
				wss.Hub.PushMsg(data)
			}
		case <-ctx.Done():
			break
		}
	}
}

func handleMsg(msg []byte) {
	fmt.Println(string(msg[:]))
}

func StopWatch() {
	cancel()
}

func NewAndWatch(s *api.WsServer) {
	wss = s
	wss.Hub.NotificationRegister(handleMsg)
	C.init_shm()

	ctx_shm, cancel := context.WithCancel(context.Background())
	go watcher(ctx)
	return
}