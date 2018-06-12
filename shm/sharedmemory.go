package shm

// #cgo CFLAGS: -I/root/mcserver/include/
// #cgo LDFLAGS: -lshare
// #include<stdlib.h>
// #include<sharedmemory.h>
import "C"
import(
	"context"
	"fmt"
	"time"
	"sync"
	"elibot-apiserver/api"
	Log "elibot-apiserver/log"
)

const (
	duration = 5
	watchPeriod = time.Second * duration
)

type FetchFunc func(chan []byte)[]byte
var watchfuncs []FetchFunc
var ModifyChan chan []byte

func REGISTERFUNC(f FetchFunc) {
	watchfuncs = append(watchfuncs, f)
}

func initWatchFuncs() {
	watchfuncs = make([]FetchFunc, 0)
	REGISTERFUNC(getPressResetIfModified)
	REGISTERFUNC(getZeroEncodeIfModified)
}

func getPressResetIfModified(modified chan []byte) []byte {
	value := C.get_press_reset()
	fmt.Println("get Press Reset ", value)
	res := fmt.Sprint(uint64(value))
	return []byte(res)
}

func getZeroEncodeIfModified(modified chan []byte) []byte {
	value := C.get_zero_encode(0)
	fmt.Println("get zero encode: ", value)
	res := fmt.Sprint(uint64(value))
	return []byte(res)
}

func fetchAndCompare(fetch FetchFunc, modified chan []byte, old []byte) {
	now := fetch()
	if now != old {
		modified <- now
	}
	return now
}

func worker(ctx_base context.Context, modified chan []byte) {
	var wg sync.WatiGroup
	for _,f := range watchfuncs {
		wg.Add(1)
		defer wg.Done()

		go func(wf FetchFunc, ctx context.Context, modified chan []byte){
			watchTicker := time.NewTicker(watchPeriod)
			defer func() {
				watchTicker.Stop()
			}()

			old := wf()
			for {
				select {
				case <-ctx.Done():
					return
				case <-watchTicker.C:
					old = fetchAndCompare(wf, modified, old)
				}
			}
		}(f, ctx_base, modified)
	}
	wg.Wait()
	Log.Info("quit for some reason")
}