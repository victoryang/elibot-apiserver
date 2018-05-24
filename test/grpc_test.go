package test

import (
        "testing"
        "log"
        "time"
        "net/http"

        "golang.org/x/net/context"
        "google.golang.org/grpc"
        hello "elibot-apiserver/serverpb"
)

const (
        address     = "localhost:2500"
)

func Test_Hello(t *testing.T) {
        // Set up a connection to the server.
        conn, err := grpc.Dial(address, grpc.WithInsecure())
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
        go func() {
                start := time.Now()
                resp, err := http.Get(address + "/v1/testSocket")
                if err!=nil {
                        t.Error("test failed")
                } else {
                        t.Log("test pass")
                        d := time.Since(start)
                        fmt.Println(resp)
                        fmt.Println("time: ", d)
                }
        }()

         // Set up a connection to the server.
        conn, err := grpc.Dial(address, grpc.WithInsecure())
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