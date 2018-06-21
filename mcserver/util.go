package mcserver

import (
    "strings"

    Log "elibot-apiserver/log"
)

/* A safe send*/
func SafeSendResponseToChannel(ch chan Response, resp Response) {
    defer func() {
        if err := recover(); err!=nil {
            Log.Error("Client closed response channel... drop it")
        }
    }()
    ch <- resp
}

func Deprecated_SafeSendResponseToChannel(ch chan Response, resp Response) {
    if _, ok := <- ch; ok {
        ch <- resp
    } else {
        Log.Error("Client closed response channel... drop it")
    }
}