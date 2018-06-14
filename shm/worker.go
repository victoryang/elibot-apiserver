package shm

// #cgo CFLAGS: -I/root/mcserver/include/
// #cgo LDFLAGS: -lshare
// #include <stdlib.h>
// #include <workeresource.h>
import "C"
import(
	"context"
	"fmt"
	"time"
	"sync"

	Log "elibot-apiserver/log"
)

const (
	duration = 5
	watchPeriod = time.Second * duration
)

type WatchFunc func(chan []byte)
var watchfuncs map[string]WatchFunc

func REGISTERFUNC(k string, f WatchFunc) {
	watchfuncs[k] = f
}

func initWatchFuncs() {
	watchfuncs = make(map[string]WatchFunc)

	REGISTERFUNC("test", testwatch)
	REGISTERFUNC("sharedResource", watchSharedResource)
	REGISTERFUNC("NV", watchNV)
}

func testwatch(modified chan []byte) {
	value := C.watch_test()
	modified<-[]byte(fmt.Println("watch test ", value))
}

func worker(ctx_base context.Context, modified chan []byte) {
	var wg sync.WaitGroup
	for _,f := range watchfuncs {
		wg.Add(1)
		defer wg.Done()

		go func(watchfunc WatchFunc, ctx context.Context, modified chan []byte){
			watchTicker := time.NewTicker(watchPeriod)
			defer func() {
				watchTicker.Stop()
			}()

			for {
				select {
				case <-ctx.Done():
					return
				case <-watchTicker.C:
					watchfunc(modified)
				}
			}
		}(f, ctx_base, modified)
	}
	wg.Wait()
	Log.Info("quit for some reason")
}

func initWorkerResource() {
	C.init_worker_resource()
	InitSharedResource()
	InitNVRam()
}