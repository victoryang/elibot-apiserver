#!/bin/sh

BUILD_DIR=$HOME/gopath/src/github.com/victoryang/elibot-apiservers
echo "Start to build elibot-apiserver"
go build -o elibot-apiserver
echo "End to build..."