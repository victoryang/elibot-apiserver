#!/bin/sh

export GOROOT=/usr/local/go
export GOPATH=/root/go
export PATH=${GOPATH}/bin:${GOROOT}/bin:${PATH}
export GO_VERSION=1.10
export GO_DOWNLOAD_URL=https://storage.googleapis.com/golang
export APISERVER_GIT=https://github.com/victoryang/elibot-apiserver.git
export MCSERVER=/root/mcserver

# install gcc-arm tool chain
sudo apt-get update >/dev/null
sudo apt-get install -y cd
sudo apt-get install -y libc6-armel-cross libc6-dev-armel-cross >/dev/null
sudo apt-get install -y binutils-arm-linux-gnueabi binutils-arm-linux-gnueabihf >/dev/null
sudo apt-get install -y libncurses5-dev >/dev/null
sudo apt-get install -y gcc-arm-linux-gnueabihf >/dev/null
sudo apt-get install -y g++-arm-linux-gnueabihf >/dev/null