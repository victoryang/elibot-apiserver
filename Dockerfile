FROM ubuntu:14.04

RUN apt-get update && apt-get install -y git wget make

RUN apt-get install -y libc6-armel-cross libc6-dev-armel-cross \
    && apt-get install -y binutils-arm-linux-gnueabi \
	&& apt-get install -y libncurses5-dev \
	&& apt-get install -y gcc-arm-linux-gnueabihf \
	&& apt-get install -y g++-arm-linux-gnueabihf

ENV GOROOT /usr/local/go
ENV GOPATH /root/go
ENV PATH ${GOPATH}/bin:${GOROOT}/bin:${PATH}
ENV GO_VERSION 1.10
ENV GO_DOWNLOAD_URL https://dl.google.com/go
RUN rm -rf ${GOROOT} \
  && wget ${GO_DOWNLOAD_URL}/go${GO_VERSION}.linux-amd64.tar.gz | tar -v -C /usr/local/ -xz \
  && mkdir -p ${GOPATH}/src ${GOPATH}/bin \
  && go version

WORKDIR ${GOPATH}/src/elibot-apiserver/

CMD ["./post-install.sh"]