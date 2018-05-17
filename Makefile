GO := /usr/bin/go
BUILD_PATH = $(PWD)/build
TEST_PATH = $(PWD)/test

# build all
all: libsqlitedb.so elibot-server.yaml
	@[ -d $(BUILD_PATH) ] || mkdir -p $(BUILD_PATH)
	@$(GO) build -o $(BUILD_PATH)/elibot-apiserver
	cp -f $(BUILD_PATH)/elibot-apiserver /usr/bin

libsqlitedb.so:
	cp $(PWD)/package/libsqlitedb.so /usr/lib/

elibot-server.yaml:
	cp $(PWD)/conf/elibot-server.yaml /etc/

test:
	echo "Start to test..."
	@$(GO) test $(TEST_PATH)/query_test.go

clean:
	rm -rf $(BUILD_PATH)