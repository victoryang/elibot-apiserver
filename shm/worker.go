package shm

// #cgo CFLAGS: -I/root/mcserver/include/
// #cgo LDFLAGS: -lrobresource
import "C"
import(
	"context"
	"errors"
	"strconv"
	"time"
	"sync"

	Log "elibot-apiserver/log"
)

const (
	duration = 100
	watchPeriod = time.Millisecond * duration
)

type WatchFunc func(chan []byte)
var watchfuncs map[string]WatchFunc

func REGISTERFUNC(k string, f WatchFunc) {
	watchfuncs[k] = f
}

func initWatchFuncs() {
	watchfuncs = make(map[string]WatchFunc)

	REGISTERFUNC("plc", watchPLC)
	REGISTERFUNC("sharedResource", watchSharedResource)
	REGISTERFUNC("NV", watchNV)
}

func worker(ctx_base context.Context, modified chan []byte) {
	var wg sync.WaitGroup
	for _,f := range watchfuncs {
		wg.Add(1)
		
		go func(watchfunc WatchFunc, ctx context.Context, modified chan []byte){
			defer wg.Done()
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
	Log.Debug("Shared memory:  workers quit")
}

func initWorkerResource() error {
	res := InitSharedResource()
	if res != 0 {
		errMsg := string("failed to init shared resource, return value:") + strconv.Itoa(int(res))
		return errors.New(errMsg)
	}
	res = InitNVRam()
	if res != 0 {
		errMsg := string("failed to init NV ram, return value:") + strconv.Itoa(int(res))
		return errors.New(errMsg)
	}
	return nil
}
