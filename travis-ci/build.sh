#!/bin/sh

sudo mkdir ${GOPATH}/src/elibot-apiserver
sudo cp -a * ${GOPATH}/src/elibot-apiserver

GO111MODULE=on CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 go build -v -o elibot-server