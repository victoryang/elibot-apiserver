package resource

import(
    "context"
    "fmt"
    "errors"
    "strconv"
    "time"
    "sync"

    "elibot-apiserver/websocket"
    Log "elibot-apiserver/log"
)

const (
    BufferSize = 1024
    Duration = 100
    WatchPeriod = time.Millisecond * Duration
)

var hit chan []byte

type WatchFunc func(chan []byte)
var watchfuncs map[string]WatchFunc
func REGISTERFUNC(k string, f WatchFunc) {
    watchfuncs[k] = f
}

type ResServer struct {
    Hit         chan []byte
    Ctx         context.Context
    Cancel      context.CancelFunc
}

func sender(ctx context.Context, hit chan []byte) {
    Log.Print("Shared memory server started...")

    DONE:
    for {
        select {
        case <-ctx.Done():
            break DONE
        case res := <-hit:
            websocket.PushBytes(res)
        }
    }
}

func worker(ctx_base context.Context, modified chan []byte) {
    var wg sync.WaitGroup
    for _,f := range watchfuncs {
        wg.Add(1)
        
        go func(watchfunc WatchFunc, ctx context.Context, modified chan []byte){
            defer wg.Done()
            watchTicker := time.NewTicker(WatchPeriod)
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

func (s *ResServer) Shutdown() {
    s.Cancel()
    Log.Print("shared memory server shutting down")
}

func (s *ResServer) Start() {
    go worker(s.Ctx, s.Hit)
    go sender(s.Ctx, s.Hit)
}

func NewServer() (*ResServer, error){
    s := new(ResServer)
    ctx, cancel := context.WithCancel(context.Background())
    s = &ResServer{
        Hit:    make(chan []byte, BufferSize),
        Ctx:    ctx,
        Cancel: cancel,
    }
    err := initWorkerResource()
    if err != nil {
        return nil, err
    }
    initWatchFuncs()
    websocket.NotificationRegister(handleMsg)
    return s, nil
}

func initWorkerResource() error {
    res := InitResource()
    if res != 0 {
        errMsg := string("failed to init shared resource, return value:") + strconv.Itoa(int(res))
        return errors.New(errMsg)
    }

    return nil
}

func initWatchFuncs() {
    watchfuncs = make(map[string]WatchFunc)

    REGISTERFUNC("plc", watchPLC)
    REGISTERFUNC("sharedResource", watchResource)
    REGISTERFUNC("NV", watchNV)
}

func handleMsg(msg []byte) {
    fmt.Println(string(msg[:]))
}