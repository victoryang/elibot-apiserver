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

# build pre-requisites for linux(automake libtool,etc...)
sudo apt-get install build-essential autoconf libtool pkg-config
# if build from source
sudo apt-get install libgflags-dev libgtest-dev
sudo apt-get install clang libc++-dev

# IMPORTANT: 
# Cross-compile for arm, see issue on https://github.com/grpc/grpc/issues/9719
git clone -b $(curl -L https://grpc.io/release) https://github.com/grpc/grpc
cd grpc
git submodule update --init

# Protobuf installation [ HOST ]
cd grpc/third_party/protobuf
./autogen.sh
./configure
make && make check && sudo make install && sudo ldconfig
cp /usr/local/bin/protoc /usr/bin/protoc

# Grpc installation [ HOST ]
cd grpc
make && sudo make install && sudo ldconfig

# clean up this build
make clean

# Grpc compile for ARM (Clean build)
# Important: make plugins will generate libaddress_sorting.a and libares.a
# And `make plugins` will be called during the `make`, so commented it
# cd grpc
# make plugins

export GRPC_CROSS_COMPILE=true
export GRPC_CROSS_AROPTS="cr --target=elf32-little"

# here comes cross-compile
make HAS_PKG_CONFIG=false \ 
CC=arm-linux-gnueabihf-gcc \
CXX=arm-linux-gnueabihf-g++ \
RANLIB=arm-linux-gnueabihf-ranlib \
LD=arm-linux-gnueabihf-ld \
LDXX=arm-linux-gnueabihf-g++ \
AR=arm-linux-gnueabihf-ar \
PROTOBUF_CONFIG_OPTS="--host=arm-linux-gnueabihf --with-protoc=/usr/local/bin/protoc" static



# PREQUISITE
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

# if we want to run example of pb, cd to example and run make to generate addressbook.pb.cc and addressbook.pb.h with protoc
protoc --cpp_out=. addressbook.proto

# configure to compile source to a arm libaray, and make 
# Hints: if we meet problem with "/usr/lib/libstdc++.so.6: version `GLIBCXX_3.4.21' not found", 
# solution 1.re-compile everything we've got with -static-libstdc++ if we want to add this option to makefile, see solutions in: https://www.zhihu.com/question/22940048
# solution 2(recommanded): choose Ubuntu 14.04 as our build platform

./configure --host=arm-linux-gnueabihf CC=arm-linux-gnueabihf-gcc CXX=arm-linux-gnueabihf-g++ --with-protoc=/usr/bin/protoc 
make
make install

# how to test examples with protobuf
arm-linux-gnueabihf-g++ add_person.cc addressbook.pb.cc -o add_person_cpp -pthread -I/usr/local/include -lprotobuf 
arm-linux-gnueabihf-g++ list_people.cc addressbook.pb.cc -o list_people_cpp -pthread -I/usr/local/include -lprotobuf 
scp libprotobuf.so.xxxx to root@${arm board ip}:/usr/lib/
scp add_person_cpp list_people_cpp root@${arm board ip}:~


# compile libares.a
CC=arm-linux-gnueabihf-gcc LD=arm-linux-gnueabihf-ld AR=arm-linux-gnueabihf-ar ./configure --host=arm-linux --disable-shared
make
for i in libcares_la-*.o; do mv $i ${i#libcares_la-}; done
arm-linux-gnueabihf-ar cr --target=elf32-little libares.a *.o
