package test

import (
        "testing"
        "log"
        "time"
        "fmt"
        "sync"
        "net/http"
        "io/ioutil"

        "golang.org/x/net/context"
        "google.golang.org/grpc"
        hello "elibot-apiserver/serverpb"
)

const (
        addressForGRPC = "192.168.1.106:2500"
        addressForAPI  = "http://192.168.1.106:9000"

        maxThreadnum = 3
)

func Test_Hello(t *testing.T) {
        // Set up a connection to the server.
        conn, err := grpc.Dial(addressForGRPC, grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()
        c := hello.NewGreeterClient(conn)

        // Contact the server and print out its response.
        var name int32 = 1
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        r, err := c.SayHello(ctx, &hello.Req{Name: name})
        if err != nil {
                log.Fatalf("could not greet: %v", err)
        }
        log.Printf("Greeting: %s", r.Message)
}

func Test_Hello_In_Para(t *testing.T) {
        var wg *sync.WaitGroup
        for i:=0;i<maxThreadnum;i++ {
               go func(i int) {
                        wg.Add(1)
                        defer wg.Done()
                        fmt.Println("In thread NO.", i)
                        start := time.Now()
                        resp, err := http.Get(addressForAPI + "/v1/testSocket")
                        if err!=nil {
                                t.Error("test failed: ", err)
                        } else {
                                t.Log("test pass")
                                d := time.Since(start)
                                buf, _ := ioutil.ReadAll(resp.Body)
                                fmt.Println("api server returns: ",string(buf))
                                fmt.Println("time: ", d)
                        }
                }(i)
        }


         // Set up a connection to the server.
        conn, err := grpc.Dial(addressForGRPC, grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()
        c := hello.NewGreeterClient(conn)

        // Contact the server and print out its response.
        var name int32 = 1
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        r, err := c.SayHello(ctx, &hello.Req{Name: name})
        if err != nil {
                log.Fatalf("could not greet: %v", err)
        }
        log.Printf("Greeting: %s", r.Message)

        wg.Wait()
}