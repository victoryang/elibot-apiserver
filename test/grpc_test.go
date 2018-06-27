package test

import (
        "testing"
        "log"
        "time"
        "fmt"
        "sync"
        "net/http"
        "io/ioutil"
        "crypto/tls"
        "crypto/x509"

        "golang.org/x/net/context"
        "google.golang.org/grpc"
        hello "elibot-apiserver/serverpb"
        "google.golang.org/grpc/credentials"
)

const (
        addressForGRPC = "192.168.1.217:2500"
        addressForAPI  = "http://192.168.1.217:9000"

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
        var wg sync.WaitGroup
        for i:=0;i<maxThreadnum;i++ {
               go func(i int) {
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
                wg.Add(1)
        }


         // Set up a connection to the server.
        start := time.Now()
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
        d := time.Since(start)
        log.Printf("Greeting: %s", r.Message)
        fmt.Println("time: ", d)

        wg.Wait()
}

func Test_Secure_Hello(t *testing.T) {
        // Set up a connection to the server.
         // load peer cert/key, cacert
        certfile := "client/client-cert.pem"
        keyfile := "client/client-key.pem"
        cafile := "ca/ca-cert.pem"
        peerCert, err := tls.LoadX509KeyPair(certfile, keyfile)
        if err != nil {
            log.Fatalf("load peer cert/key error:%v", err)
        }
        caCert, err := ioutil.ReadFile(cafile)
        if err != nil {
            log.Fatalf("read ca cert file error:%v", err)
        }
        caCertPool := x509.NewCertPool()
        caCertPool.AppendCertsFromPEM(caCert)

        ta := credentials.NewTLS(&tls.Config{
            Certificates: []tls.Certificate{peerCert},
            RootCAs:      caCertPool,
        })
        conn, err := grpc.Dial(addressForGRPC, grpc.WithTransportCredentials(ta))
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