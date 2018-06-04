#!/bin/sh

# To cross-compile for our arm on ubuntu, we need arm-linux-gnueabihf-gcc tools(ld-linux-armhf.so.x installed on arm board)
sudo apt-get install -y libc6-armel-cross libc6-dev-armel-cross
sudo apt-get install -y binutils-arm-linux-gnueabi binutils-arm-linux-gnueabihf
sudo apt-get install -y libncurses5-dev

# For different OS:  SYSV interpreted by /lib/ld-linux-armhf.so.x
sudo apt-get install -y gcc-arm-linux-gnueabihf
sudo apt-get install -y g++-arm-linux-gnueabihf
# Or GNU/LINUX interpreted by /lib/ld-linux.so.x
sudo apt-get install -y gcc-arm-linux-gnueabi
sudo apt-get install -y g++-arm-linux-gnueabi

# Install libz.so.1 for dependency of libsqlitedb.so
#sudo apt-get install lib64z1 
# This library should be replaced with the one which is competible to libsqlitedb.so
# which is provided in build directory(a linux sysV version for arm board)
cp libsqlitedb /usr/lib/gcc-cross/arm-linux-gnueabihf/5/
cp libz.so.1.2.8 /usr/arm-linux-gnueabihf/lib/ 
ln -s /usr/arm-linux-gnueabihf/lib/libz.so.1.2.8 /usr/arm-linux-gnueabihf/lib/libz.so.1

#env GOOS=linux GOARCH=arm64 go build -o elibot-server

CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 go build -v -o /root/elibot-server