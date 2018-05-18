#!/bin/sh

# To cross compile on ubuntu, we need arm-linux-gnueabihf-gcc tools
sudo apt-get install libc6-armel-cross libc6-dev-armel-cross
sudo apt-get install binutils-arm-linux-gnueabi
sudo apt-get install libncurses5-dev

# For different board
sudo apt-get install gcc-arm-linux-gnueabihf
sudo apt-get install g++-arm-linux-gnueabihf
# Or
sudo apt-get install gcc-arm-linux-gnueabi
sudo apt-get install g++-arm-linux-gnueabi

# Install libz.so.1 for dependency of libsqlitedb.so
sudo apt-get install lib64z1
ln -s /usr/arm-linux-gnueabi/lib/libz.so.1.2.8 /usr/lib/gcc-cross/arm-linux-gnueabihf/5/../../../../arm-linux-gnueabihf/lib/libz.so.1

#env GOOS=linux GOARCH=arm64 go build -o elibot-server

CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 go build -v -o /root/elibot-server #-ldflags="-extld=$CC"