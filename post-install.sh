#!/bin/sh

cp build/lib/libsqlitedb.so /usr/lib/gcc-cross/arm-linux-gnueabihf/4.8/
cp build/lib/libz.so.1.2.8 /usr/arm-linux-gnueabihf/lib/ 
ln -s /usr/arm-linux-gnueabihf/lib/libz.so.1.2.8 /usr/arm-linux-gnueabihf/lib/libz.so.1
cp build/lib/librobresource.so /usr/arm-linux-gnueabihf/lib/

/bin/bash