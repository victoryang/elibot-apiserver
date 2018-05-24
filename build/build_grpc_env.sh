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



# install grpc from source
# see it on https://github.com/grpc/grpc/blob/master/INSTALL.md

# build pre-requisites for linux
apt-get install build-essential autoconf libtool pkg-config
# if build from source
apt-get install libgflags-dev libgtest-dev
apt-get install clang libc++-dev

# Cross-compile for arm, see issue on https://github.com/grpc/grpc/issues/9719



