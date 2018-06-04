FROM ubuntu:14.04

RUN apt-get update && apt-get install -y git curl

RUN apt-get install -y libc6-armel-cross libc6-dev-armel-cross \
    && apt-get install -y binutils-arm-linux-gnueabi \
	&& apt-get install -y libncurses5-dev \
	&& apt-get install -y gcc-arm-linux-gnueabihf \
	&& apt-get install -y g++-arm-linux-gnueabihf

RUN apt-get install build-essential autoconf libtool pkg-config \
	&& apt-get install libgflags-dev libgtest-dev \
	&& apt-get install clang libc++-dev \

ENV APISERVER_GIT https://github.com/victoryang/elibot-apiserver.git
RUN mkdir -p ${GOPATH}/src/ \
	&& cd ${GOPATH}/src/ \
	&& git clone  ${APISERVER_GIT} \
	&& cd elibot-apiserver/ \
	&& cp build/libz.so.1.2.8  /usr/arm-linux-gnueabihf/lib/ \
	&& ln -s /usr/arm-linux-gnueabihf/lib/libz.so.1.2.8 /usr/arm-linux-gnueabihf/lib/libz.so.1 \

RUN cd ~ \
	&& git clone -b $(curl -L https://grpc.io/release) https://github.com/grpc/grpc \
	&& cd grpc \
	&& git submodule update --init \
	&& cd third-party/protobuf \
	&& ./autogen.sh && ./configure \
	&& make && make check && sudo make install && sudo ldconfig