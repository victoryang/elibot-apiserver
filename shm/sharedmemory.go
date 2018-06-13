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
	"reflect"
	"runtime"

	Log "elibot-apiserver/log"
)

const (
	duration = 5
	watchPeriod = time.Second * duration
)

type FetchFunc func()string
var watchfuncs map[string]FetchFunc

func REGISTERFUNC(k string, f FetchFunc) {
	watchfuncs[k] = f
}

func initWatchFuncs() {
	watchfuncs = make(map[string]FetchFunc)

	REGISTERFUNC("test", testwatch)
	REGISTERFUNC("NV", watchNV)
}

func testwatch() string {
	value := C.watch_test()
	fmt.Println("watch test ", value)
	return fmt.Sprint(uint64(value))
}

func GetFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func fetchAndCompare(key string, modified chan []byte, old string) string{
	now := watchfuncs[key]()
	if now != old {
		modified <- byte[now]
	}
	return now
}

func worker(ctx_base context.Context, modified chan []byte) {
	var wg sync.WaitGroup
	for k,_ := range watchfuncs {
		wg.Add(1)
		defer wg.Done()

		go func(key string, ctx context.Context, modified chan []byte){
			watchTicker := time.NewTicker(watchPeriod)
			defer func() {
				watchTicker.Stop()
			}()

			now := watchfuncs[key]()
			modified <- []byte(now)
			for {
				select {
				case <-ctx.Done():
					return
				case <-watchTicker.C:
					now = fetchAndCompare(key, modified, now)
				}
			}
		}(k, ctx_base, modified)
	}
	wg.Wait()
	Log.Info("quit for some reason")
}