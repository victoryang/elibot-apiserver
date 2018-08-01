#!/bin/sh

cp lib/libsqlitedb /usr/lib/gcc-cross/arm-linux-gnueabihf/5/
cp lib/libz.so.1.2.8 /usr/arm-linux-gnueabihf/lib/ 
ln -s /usr/arm-linux-gnueabihf/lib/libz.so.1.2.8 /usr/arm-linux-gnueabihf/lib/libz.so.1
cp lib/libshare.a /usr/arm-linux-gnueabihf/lib/
mkdir /root/mcserver
tar -xvf include/include.tar.gz -C /root/mcserver/
cp include/config.h /root/mcserver/