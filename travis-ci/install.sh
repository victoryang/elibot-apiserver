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
sudo apt-get install -y git curl
sudo apt-get install -y libc6-armel-cross libc6-dev-armel-cross
sudo apt-get install -y binutils-arm-linux-gnueabi binutils-arm-linux-gnueabihf
sudo apt-get install -y libncurses5-dev
sudo apt-get install -y gcc-arm-linux-gnueabihf
sudo apt-get install -y g++-arm-linux-gnueabihf

# include go tool
sudo mkdir ${GOROOT}
sudo mkdir ${GOPATH}
sudo curl -s ${GO_DOWNLOAD_URL}/go${GO_VERSION}.linux-amd64.tar.gz 
sudo tar -v -C /usr/local/ -xz go${GO_VERSION}.linux-amd64.tar.gz
sudo mkdir -p ${GOPATH}/src ${GOPATH}/bin

# install build dependency
sudo cd ${GOPATH}/src/
sudo git clone ${APISERVER_GIT}
sudo cd elibot-apiserver/build/
sudo cp lib/libsqlitedb.so /usr/lib/gcc-cross/arm-linux-gnueabihf/4.8
sudo cp lib/libz.so.1.2.8  /usr/arm-linux-gnueabihf/lib/
sudo ln -s /usr/arm-linux-gnueabihf/lib/libz.so.1.2.8 /usr/arm-linux-gnueabihf/lib/libz.so.1

sudo mkdir MCSERVER
sudo tar -xvf include/include.tar.gz -C ${MCSERVER}
sudo cp lib/libshare.a /usr/arm-linux-gnueabihf/lib/