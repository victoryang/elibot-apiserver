FROM ubuntu:14.04

RUN apt-get update && apt-get install -y git curl unzip

RUN apt-get install -y libc6-armel-cross libc6-dev-armel-cross \
    && apt-get install -y binutils-arm-linux-gnueabi \
	&& apt-get install -y libncurses5-dev \
	&& apt-get install -y gcc-arm-linux-gnueabihf \
	&& apt-get install -y g++-arm-linux-gnueabihf

RUN apt-get install -y build-essential autoconf libtool pkg-config \
	&& apt-get install -y libgflags-dev libgtest-dev \
	&& apt-get install -y clang libc++-dev

ENV GOPATH /root/go
ENV APISERVER_GIT https://github.com/victoryang/elibot-apiserver.git
RUN mkdir -p ${GOPATH}/src/ \
	&& cd ${GOPATH}/src/ \
	&& git clone ${APISERVER_GIT}

RUN cd ~ \
	&& git clone -b $(curl -L https://grpc.io/release) https://github.com/grpc/grpc \
	&& cd grpc \
	&& git submodule update --init \
	&& cd third_party/protobuf \
	&& ./autogen.sh && ./configure \
	&& make && make check && sudo make install && sudo ldconfig \
	&& cp /usr/local/bin/protoc /usr/bin/protoc \
	&& echo "enter ~/grpc" \
	&& cd ~/grpc \
	&& make && sudo make install && sudo ldconfig \
	&& make clean \
	&& export GRPC_CROSS_COMPILE=true \
	&& export GRPC_CROSS_AROPTS="cr --target=elf32-little" \
	&& make HAS_PKG_CONFIG=false CC=arm-linux-gnueabihf-gcc CXX=arm-linux-gnueabihf-g++ \
		RANLIB=arm-linux-gnueabihf-ranlib LD=arm-linux-gnueabihf-ld LDXX=arm-linux-gnueabihf-g++ AR=arm-linux-gnueabihf-ar \
		PROTOBUF_CONFIG_OPTS="--host=arm-linux-gnueabihf --with-protoc=/usr/local/bin/protoc" static \
	&& mkdir -p /usr/local/cross/lib \
	&& cp -a libs/opt/* /usr/local/cross/lib

WORKDIR ${GOPATH}/src/elibot-apiserver/