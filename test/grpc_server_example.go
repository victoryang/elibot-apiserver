package main

import (
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "io/ioutil"
    "net/http"
)

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter,
                   r *http.Request) {
    fmt.Fprintf(w,
        "Hi, This is an example of http service in golang!\n")
}

func main() {
    pool := x509.NewCertPool()
    caCertPath := "ca/ca-cert.pem"

    caCrt, err := ioutil.ReadFile(caCertPath)
    if err != nil {
        fmt.Println("ReadFile err:", err)
        return
    }
    pool.AppendCertsFromPEM(caCrt)

    s := &http.Server{
        Addr:    ":8081",
        Handler: &myhandler{},
        TLSConfig: &tls.Config{
            ClientCAs:  pool,
            ClientAuth: tls.RequireAndVerifyClientCert,
        },
    }

    err = s.ListenAndServeTLS("server/server-cert.pem", "server/server-key.pem")
    if err != nil {
        fmt.Println("ListenAndServeTLS err:", err)
    }
}