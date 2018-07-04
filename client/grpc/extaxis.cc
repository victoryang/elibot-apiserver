#include <string>
#include <grpcpp/grpcpp.h>

#include "rpc.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;

using rpc::Req;
using rpc::Reply;
using rpc::ExtAxis;

class GrpcClient {
public:
	GrpcClient(std::shared_ptr<Channel> channel)
		: stub_(ExtAxis::NewStub(channel)) {}

	std::string GotoExtaxisPos(::google::protobuf::int32& axis, ::google::protobuf::int32& num) {
		Req request;
	    request.set_axis(axis);
	    request.set_num(num);

	    // Container for the data we expect from the server.
	    Reply reply;

	    // Context for the client. It could be used to convey extra information to
	    // the server and/or tweak certain RPC behaviors.
	    ClientContext context;

	    // The actual RPC.
	    Status status = stub_->GotoExtaxisPos(&context, request, &reply);

	    // Act upon its status.
	    if (status.ok()) {
	      return reply.res();
	    } else {
	      std::cout << status.error_code() << ": " << status.error_message()
	                << std::endl;
	      return "RPC failed";
    	}
	}

private:
	std::unique_ptr<ExtAxis::Stub> stub_;
};

int main(int argc, char** argv) {
  // Instantiate the client. It requires a channel, out of which the actual RPCs
  // are created. This channel models a connection to an endpoint (in this case,
  // localhost at port 9500). We indicate that the channel isn't authenticated
  // (use of InsecureChannelCredentials()).
  GrpcClient client(grpc::CreateChannel(
      "localhost:9500", grpc::InsecureChannelCredentials()));

  ::google::protobuf::int32 i = 1;
  ::google::protobuf::int32 j = 1;
  std::string reply = client.GotoExtaxisPos(i, j);
  std::cout << "Greeter received: " << reply << std::endl;

  return 0;
}