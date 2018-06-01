CXX = arm-linux-gnueabihf-g++
CPPFLAGS += -I/usr/local/include -pthread -static
CXXFLAGS += -std=c++11
LDFLAGS  += -L/root/grpc_for_arm_cross_compile/lib -L/root/grpc_for_arm_cross_compile/lib/protobuf \
		   -lgrpc++ -lgrpc    \
           -lgrpc++_reflection \
           -lprotobuf -lpthread -ldl

PROTOC = protoc
GRPC_CPP_PLUGIN = grpc_cpp_plugin
GRPC_CPP_PLUGIN_PATH ?= `which $(GRPC_CPP_PLUGIN)`
PROTOS_PATH = ../serverpb

vpath %.proto $(PROTOS_PATH)

.PHONY: all

all: hello_client

hello_client: hello.pb.o hello.grpc.pb.o hello_client.o
	$(CXX) $^ $(LDFLAGS) -lares -o $@

.PRECIOUS: %.grpc.pb.cc
%.grpc.pb.cc: %.proto
	$(PROTOC) -I $(PROTOS_PATH) --grpc_out=. --plugin=protoc-gen-grpc=$(GRPC_CPP_PLUGIN_PATH) $<

.PRECIOUS: %.pb.cc
%.pb.cc: %.proto
	$(PROTOC) -I $(PROTOS_PATH) --cpp_out=. $<

clean:
	rm -f *.o *.pb.cc *.pb.h hello_client