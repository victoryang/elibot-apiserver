GO = /usr/bin/go
BUILD_PATH = $(PWD)/build
TEST_PATH = $(PWD)/test
KILLALL = /usr/bin/killall
SQLITEDB_DIR = /usr/lib/
LIBSQLITEDB = /usr/lib/libsqlitedb.so

# build all
all: libsqlitedb.so elibot-server.yaml
	@[ -d $(BUILD_PATH) ] || mkdir -p $(BUILD_PATH)
	@$(GO) build -o $(BUILD_PATH)/elibot-apiserver
	cp -f $(BUILD_PATH)/elibot-apiserver /usr/bin

libsqlitedb.so:
	@if [ ! -f $(LIBSQLITEDB) ]; then \
		 echo "Cannot find libsqlitedb.so in ${SQLITEDB_DIR}, please install"; \
		 echo "fail to build..."; \
		 exit -1; \
	 fi

elibot-server.yaml:
	cp $(PWD)/conf/elibot-server.yaml /etc/

test: all
	echo "Start to test..."
	elibot-apiserver& > /dev/null
	@$(GO) test $(TEST_PATH)/query_test.go

clean:
	rm -rf $(BUILD_PATH)
	@$(KILLALL) -9 elibot-apiserver