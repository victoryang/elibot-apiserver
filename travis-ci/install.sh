#!/bin/sh

# install gcc-arm tool chain
sudo apt-get update >/dev/null
sudo apt-get install -y cd
sudo apt-get install -y libc6-armel-cross libc6-dev-armel-cross >/dev/null
sudo apt-get install -y binutils-arm-linux-gnueabi binutils-arm-linux-gnueabihf >/dev/null
sudo apt-get install -y libncurses5-dev >/dev/null
sudo apt-get install -y gcc-arm-linux-gnueabihf >/dev/null
sudo apt-get install -y g++-arm-linux-gnueabihf >/dev/null

sudo cp build/lib/libsqlitedb.so /usr/lib/gcc-cross/arm-linux-gnueabihf/4.8
sudo cp build/lib/libz.so.1.2.8  /usr/arm-linux-gnueabihf/lib/
sudo ln -s /usr/arm-linux-gnueabihf/lib/libz.so.1.2.8 /usr/arm-linux-gnueabihf/lib/libz.so.1
sudo mkdir /usr/local/include/mcserver/
sudo tar -xvf build/include/include.tar.gz -C /usr/local/include/mcserver/
sudo cp build/include/config.h /usr/local/include/mcserver/
sudo cp build/lib/librobresource.so /usr/arm-linux-gnueabihf/lib/