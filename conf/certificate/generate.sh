#!/bin/sh
# self-signing for creating CA
# create private key
openssl genrsa -out ca/ca-key.pem 1024

# create certificate request
openssl req -new -out ca/ca-req.csr -key ca/ca-key.pem

# create certificate
openssl x509 -req -in ca/ca-req.csr -out ca/ca-cert.pem -signkey ca/ca-key.pem -days 3650


# create server certificate
# create private key
openssl genrsa -out server/server-key.pem 1024

# request
openssl req -new -out server/server-req.csr -key server/server-key.pem

# certificate
openssl x509 -req -in server/server-req.csr -out server/server-cert.pem -signkey server/server-key.pem -CA ca/ca-cert.pem -CAkey ca/ca-key.pem -CAcreateserial -days 3650

# for client
openssl genrsa -out client/client-key.pem 1024
openssl req -new -out client/client-req.csr -key client/client-key.pem
openssl x509 -req -in client/client-req.csr -out client/client-cert.pem -signkey client/client-key.pem -CA ca/ca-cert.pem -CAkey ca/ca-key.pem -CAcreateserial -days 3650


# to verify generated cert files with ca
openssl verify -CAfile ca/ca-cert.pem client/client-cert.pem

# to verify the server with specific client pem and ca pem
openssl s_client -connect 192.168.1.217:9500 -tls1 -cert client/client-cert.pem -key client/client-key.pem -CAfile ca/ca-cert.pem

# generating tools
git clone https://github.com/square/certstrap
