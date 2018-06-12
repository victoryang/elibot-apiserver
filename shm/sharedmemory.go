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
	"strconv"
	"reflect"
	"runtime"

	Log "elibot-apiserver/log"
)

const (
	duration = 5
	watchPeriod = time.Second * duration
)

type FetchFunc func()string
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

func getPressResetIfModified() string {
	value := C.get_press_reset()
	fmt.Println("get Press Reset ", value)
	return fmt.Sprint(uint64(value))
}

func getZeroEncodeIfModified() string{
	value := C.get_zero_encode(0)
	fmt.Println("get zero encode: ", value)
	return fmt.Sprint(uint64(value))
}

func GetFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func fetchAndCompare(fetch FetchFunc, modified chan []byte, old string) string{
	now := fetch()
	now_v := strconv.Atoi(now)
	old_v := strconv.Atoi(old)
	if now_v != old_v {
		funcName := GetFunctionName(fetch)
		modified <- []byte(funcName+now)
	}
	return now
}

func worker(ctx_base context.Context, modified chan []byte) {
	var wg sync.WaitGroup
	for _,f := range watchfuncs {
		wg.Add(1)
		defer wg.Done()

		go func(wf FetchFunc, ctx context.Context, modified chan []byte){
			watchTicker := time.NewTicker(watchPeriod)
			defer func() {
				watchTicker.Stop()
			}()

			now := wf(modified)
			modified <- []byte(now)
			for {
				select {
				case <-ctx.Done():
					return
				case <-watchTicker.C:
					now = fetchAndCompare(wf, modified, now)
				}
			}
		}(f, ctx_base, modified)
	}
	wg.Wait()
	Log.Info("quit for some reason")
}