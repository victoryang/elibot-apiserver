package test

import (
	"testing"
	"time"
	"context"
	"net/http"
	"fmt"
)

cosnt (
	Intervals = 100 * time.Millisecond
	Duration = 60 * time.Second
)

func Test_Shm(t *testing.T) {
	start := time.Now()
	var success int = 0
	ctx, cancel := context.WithTimeout(context.Background(), Duration)
	defer cancel()
	ticker := time.NewTicker(Intervals)
Loop:
	for {
		select {
		case <-ticker.C:
			resp, err := http.Get("http://192.168.1.253:9000/v2/resource/sysvar/irobi?start=0&&end=30")
			if err!=nil {
                //fmt.Println("error happens")
                break
            }
			if resp.StatusCode == 200 {
				//fmt.Println("request success")
				success++
			}
		case <-ctx.Done():
			break Loop
		}
	}
	d := time.Since(start)
    fmt.Println("-----------------test finished-------------")
    fmt.Println("request url: http://192.168.1.253:9000/v2/resource/sysvar/irobi?start=0&&end=30")
    fmt.Println("Duration: ", d)
    fmt.Println("successful request number: ", success)
}