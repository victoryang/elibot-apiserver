#!/bin/sh

make -C thirdparty/mcsql/
make -C thirdparty/mrj/

CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 go build -v -o elibot-server