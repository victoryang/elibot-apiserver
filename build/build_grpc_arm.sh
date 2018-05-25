#!/bin/sh

# install grpc package for go
go get -u google.golang.org/grpc
# if you can't get that package, try clone it
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc

# download the Proto buffers binaries from github
https://github.com/google/protobuf/releases   protoc-<version>-<platform>.zip #unzip it and got protoc application

# install protobuf plugin for Go, and expose the binary
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
export PATH=$PATH:$GOPATH/bin

# generate *.pb.go file from *.proto file
protoc --go_out=plugins=grpc:. *.proto



# install grpc from source, first try protobuf in the next section
# see it on https://github.com/grpc/grpc/blob/master/INSTALL.md

# build pre-requisites for linux
apt-get install build-essential autoconf libtool pkg-config
# if build from source
apt-get install libgflags-dev libgtest-dev
apt-get install clang libc++-dev

# Cross-compile for arm, see issue on https://github.com/grpc/grpc/issues/9719
git clone -b $(curl -L https://grpc.io/release) https://github.com/grpc/grpc
cd grpc
git submodule update --init
# here comes cross-compile
CPPFLAGS=-I/usr/arm-linux-gnueabihf/include LDFLAGS=-L/usr/arm-linux-gnueabihf/lib \
make HAS_PKG_CONFIG=false CC=arm-linux-gnueabihf-gcc CXX=arm-linux-gnueabihf-g++ LD=arm-linux-gnueabihf-ld \
LDXX=arm-linux-gnueabihf-g++ AR=arm-linux-gnueabihf-ar 




# build protobuf for arm before we try Grpc
# Hints: issue for cross compile of protobuf, see https://github.com/google/protobuf/tree/master/src 
# https://github.com/google/protobuf/blob/master/src/README.md

# IMPORTANT: https://raspberrypi.stackexchange.com/questions/24448/cross-compiling-protobuf-for-raspberry-pi
# To compile protoc executable binary, we should have a copy of x86 version of protoc binary
# the one should be installed into /usr/bin/ directory
cd ~/protobuf-2.6.0
./configure --disable-shared
make
make check
sudo make install
cp /usr/local/bin/protoc /usr/bin/

# then clean the binaries we compiled in $PWD/src/.lib/*
make distclean

# configure to compile source to a arm libaray, and make 
./configure --host=arm-linux-gnueabihf CC=arm-linux-gnueabihf-gcc CXX=arm-linux-gnueabihf-g++ --with-protoc=/usr/bin/protoc 
make
make install
