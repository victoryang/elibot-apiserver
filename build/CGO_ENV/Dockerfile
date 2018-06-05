FROM ubuntu:14.04

RUN apt-get update && apt-get install -y git curl

RUN apt-get install -y libc6-armel-cross libc6-dev-armel-cross \
    && apt-get install -y binutils-arm-linux-gnueabi \
	&& apt-get install -y libncurses5-dev \
	&& apt-get install -y gcc-arm-linux-gnueabihf \
	&& apt-get install -y g++-arm-linux-gnueabihf

ENV GOROOT /usr/local/go
ENV GOPATH /root/go
ENV PATH ${GOPATH}/bin:${GOROOT}/bin:${PATH}
ENV GO_VERSION 1.10
ENV GO_DOWNLOAD_URL https://storage.googleapis.com/golang
RUN rm -rf ${GOROOT} \
  && curl -s ${GO_DOWNLOAD_URL}/go${GO_VERSION}.linux-amd64.tar.gz | tar -v -C /usr/local/ -xz \
  && mkdir -p ${GOPATH}/src ${GOPATH}/bin \
  && go version

ENV APISERVER_GIT https://github.com/victoryang/elibot-apiserver.git
RUN mkdir -p ${GOPATH}/src/ \
	&& cd ${GOPATH}/src/ \
	&& git clone  ${APISERVER_GIT} \
	&& cd elibot-apiserver/build/CGO_ENV/ \
	&& cp libsqlitedb.so /usr/lib/gcc-cross/arm-linux-gnueabihf/4.8 \
	&& cp libz.so.1.2.8  /usr/arm-linux-gnueabihf/lib/ \
	&& ln -s /usr/arm-linux-gnueabihf/lib/libz.so.1.2.8 /usr/arm-linux-gnueabihf/lib/libz.so.1

WORKDIR ${GOPATH}/src/elibot-apiserver/
CMD git pull
