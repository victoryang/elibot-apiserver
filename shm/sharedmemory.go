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
type CheckFunc func()[]byte
var watchfuncs map[string]FetchFunc
var checkfuncs map[string]CheckFunc

func REGISTERFUNC(k string, f FetchFunc, c CheckFunc) {
	watchfuncs[k] = f
	checkfuncs[k] = c
}

func initWatchFuncs() {
	watchfuncs = make(map[string]FetchFunc)
	checkfuncs = make(map[string]CheckFunc)

	REGISTERFUNC("test", testwatch, testcheck)
	REGISTERFUNC("NV", watchNV, checkNV)
}

func testcheck() []byte {
	return []byte("testcheck")
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
		modified <- checkfuncs[key]()
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