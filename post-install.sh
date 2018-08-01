#!/bin/sh

cp build/lib/libsqlitedb.so /usr/lib/gcc-cross/arm-linux-gnueabihf/4.8/
cp build/lib/libz.so.1.2.8 /usr/arm-linux-gnueabihf/lib/ 
ln -s /usr/arm-linux-gnueabihf/lib/libz.so.1.2.8 /usr/arm-linux-gnueabihf/lib/libz.so.1
cp build/lib/libshare.a /usr/arm-linux-gnueabihf/lib/
mkdir /root/mcserver
tar -xvf build/include/include.tar.gz -C /root/mcserver/ >/dev/null
cp build/include/config.h /root/mcserver/

/bin/bash