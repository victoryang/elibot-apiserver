package v2

import (
    "crypto/tls"
    "crypto/x509"
    "errors"
    "io/ioutil"
    "elibot-apiserver/auth"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
)

func LoadSelfSignedCert() (grpc.ServerOption, error){
	 // Load the certificates from disk
    certificate, err := tls.LoadX509KeyPair(auth.GetCert(), auth.GetKey())
    if err != nil {
        return nil, err
    }

    // Create a certificate pool from the certificate authority
    certPool := x509.NewCertPool()
    ca, err := ioutil.ReadFile("/rbctrl/apiserver/certificate/ca/ca-cert.pem")
    if err != nil {
        return nil, err
    }

    // Append the client certificates from the CA
    if ok := certPool.AppendCertsFromPEM(ca); !ok {
        return nil, errors.New("failed to append client certs")
    }

    // Create the TLS credentials
    cred := credentials.NewTLS(&tls.Config{
    	ClientAuth:   tls.RequireAndVerifyClientCert,
    	Certificates: []tls.Certificate{certificate},
    	ClientCAs:    certPool,
    })
    return grpc.Creds(cred), nil
}

func LoadCert() (grpc.ServerOption, error) {
    return credentials.NewServerTLSFromFile(auth.GetCert(), auth.GetKey())
}