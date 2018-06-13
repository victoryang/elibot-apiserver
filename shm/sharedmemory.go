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
var watchfuncs map[FetchFunc]CheckFunc
var ModifyChan chan []byte

func REGISTERFUNC(f FetchFunc, c CheckFunc) {
	watchfuncs[f] = c
}

func initWatchFuncs() {
	watchfuncs = make(map[FetchFunc]CheckFunc)
	REGISTERFUNC(testwatch, testcheck)
	REGISTERFUNC(watchNV, checkNV)
}

func testcheck() []byte {
	return []byte("testcheck")
}

func testwatch() string {
	value := C.watch_test()
	fmt.Println("watch test ", value)
	return fmt.Sprint(uint64(value))
}

func getZeroEncodeIfModified() string{
	value := C.get_zero_encode(0)
	fmt.Println("get zero encode: ", value)
	return fmt.Sprint(uint64(value))
}

func checkNV() []byte {
	return []byte("checkNV")
}

func watchNV() string {
	value := C.watch_nv_hash()
	return fmt.Sprint(uint64(value))
}

func GetFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func fetchAndCompare(fetch FetchFunc, cf CheckFunc, modified chan []byte, old string) string{
	now := fetch()
	if now != old {
		modified <- cf()
	}
	return now
}

func worker(ctx_base context.Context, modified chan []byte) {
	var wg sync.WaitGroup
	for f,c := range watchfuncs {
		wg.Add(1)
		defer wg.Done()

		go func(wf FetchFunc, cf CheckFunc, ctx context.Context, modified chan []byte){
			watchTicker := time.NewTicker(watchPeriod)
			defer func() {
				watchTicker.Stop()
			}()

			now := wf()
			modified <- []byte(now)
			for {
				select {
				case <-ctx.Done():
					return
				case <-watchTicker.C:
					now = fetchAndCompare(wf, cf, modified, now)
				}
			}
		}(f, c, ctx_base, modified)
	}
	wg.Wait()
	Log.Info("quit for some reason")
}