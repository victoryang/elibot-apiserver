GO = go
MAKE = make
MCSQL = thirdparty/mcsql/
MRJ = thirdparty/mrj/

.PHONY: all

all: api-server

api-server: libmrj.so libmcsql.so
	CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 $(GO) build -v -mod=vendor -o $@

clean:
	$(MAKE) -C $(MCSQL) clean
	$(MAKE) -C $(MRJ) clean
	rm api-server

libmrj.so:
	$(MAKE) -C $(MCSQL)

libmcsql.so:
	$(MAKE) -C $(MRJ)
