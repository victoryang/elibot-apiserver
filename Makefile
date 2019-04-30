GO = go
MAKE = make
MCSQL = thirdparty/mcsql/
MRJ = thirdparty/mrj/

.PHONY: all

all: api-server

api-server:
	GO111MODULE=on CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 $(GO) build -v -mod=vendor -o $@

clean:
	rm api-server
