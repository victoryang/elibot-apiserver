#include <iostream>
#include <memory>
#include <string>
#include <sstream>
#include <fstream>

#include <grpcpp/grpcpp.h>

#include "hello.grpc.pb.h"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using serverpb::Req;
using serverpb::Reply;
using serverpb::Greeter;

class GreeterClient {
 public:
  GreeterClient(std::shared_ptr<Channel> channel)
      : stub_(Greeter::NewStub(channel)) {}

  // Assembles the client's payload, sends it and presents the response back
  // from the server.
  std::string SayHello(::google::protobuf::int32& value) {
    // Data we are sending to the server.
    Req request;
    request.set_name(value);

    // Container for the data we expect from the server.
    Reply reply;

    // Context for the client. It could be used to convey extra information to
    // the server and/or tweak certain RPC behaviors.
    ClientContext context;

    // The actual RPC.
    Status status = stub_->SayHello(&context, request, &reply);

    // Act upon its status.
    if (status.ok()) {
      return reply.message();
    } else {
      std::cout << status.error_code() << ": " << status.error_message()
                << std::endl;
      return "RPC failed";
    }
  }

 private:
  std::unique_ptr<Greeter::Stub> stub_;
};

int main(int argc, char** argv) {
  // Instantiate the client channel with a composite channel, including a 
  // ssl channelcredentials as channel credential and jwt access credential
  // as call credential. It requires a channel, out of which the actual RPCs
  // are created. This channel models a connection to an endpoint (in this 
  // case, localhost at port 9500).

  /*GreeterClient greeter(grpc::CreateChannel(
      "localhost:9500", grpc::CompositeChannelCredentials(
        grpc::SslCredentials(grpc::SslCredentialsOptions()),
        grpc::ServiceAccountJWTAccessCredentials(key))));*/
  /*grpc::string ca = R"(-----BEGIN CERTIFICATE-----
MIICJzCCAZACCQColNzq5zLL5DANBgkqhkiG9w0BAQsFADBYMQswCQYDVQQGEwJB
VTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50ZXJuZXQgV2lkZ2l0
cyBQdHkgTHRkMREwDwYDVQQDDAhlbGlib3RDQTAeFw0xODA2MjYwOTIzNTlaFw0y
ODA2MjMwOTIzNTlaMFgxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRl
MSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQxETAPBgNVBAMMCGVs
aWJvdENBMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC7mcpjAX1e/Tk+QS4z
nEuXwM4wcvWMR6Ncvf1LRyQekNcMlk3a2zoqnvpijaAclMCIwx0WN7lQQIepSdwg
ec1ng+so4i594ZBsuNjAsjb1o4baWPJGHQ8v6ynSdtgv4yG3ghUzPUjdVyQ5JskK
yWCxdrvWGP5UBvIkFhLf6jvxdQIDAQABMA0GCSqGSIb3DQEBCwUAA4GBAE0dfxDQ
MkrZYGIeNhlXnyQ9HRIhxr85632dFtHpA08u2GR5x/vxic2fqNEiRGikbxD0YKKZ
cOc5ejCEkCZMOovogMScuq+8T3tiSWT29utJIALDplE2oJ6UnKID0LXdI7hXjXz8
dzzRLQr4TGhvck+UXLnfW+9fav6LB2ceMoqJ
-----END CERTIFICATE-----)";

  grpc::string cert = R"(-----BEGIN CERTIFICATE-----
MIICLDCCAZUCCQClijA1AAawYzANBgkqhkiG9w0BAQsFADBYMQswCQYDVQQGEwJB
VTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50ZXJuZXQgV2lkZ2l0
cyBQdHkgTHRkMREwDwYDVQQDDAhlbGlib3RDQTAeFw0xODA2MjcwNjE5NTlaFw0y
ODA2MjQwNjE5NTlaMF0xCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRl
MSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQxFjAUBgNVBAMMDWlt
eDZkbHdpc2VobWkwgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBAMqCAvJTq+mC
JC1uFFQPJVjg7kjKoPRq5Uke8YYlEQpvHr5stfFvVNsfN0+mhfHtoVyu59/4H87I
Z9/1Lc+JvbIGroi6G86+MNe6mbt7GIxhgUAia/2/EE0yk3en3cJ4yCTEvip3lTzy
NfmgVWoA4bAc8b+jALdssk1mYQdCmB1ZAgMBAAEwDQYJKoZIhvcNAQELBQADgYEA
rMHoULV/9G18kn42/NFNDjKyfg+PRnMjMve8xkzNOt5g1kUZfyyhT2zAhUSoSBeH
PAJGzYFCiJ7XRon4KBhX/Ymr2WXaDFY4V7eYyn4SH13/luGrdrrH+VQ7lFnRyVyo
7tjKXzh6WcH5EX2gMpQiyAOiYh8+6NxZkTS4pweW1sM=
-----END CERTIFICATE-----)";

  grpc::string key = R"(-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDKggLyU6vpgiQtbhRUDyVY4O5IyqD0auVJHvGGJREKbx6+bLXx
b1TbHzdPpoXx7aFcruff+B/OyGff9S3Pib2yBq6IuhvOvjDXupm7exiMYYFAImv9
vxBNMpN3p93CeMgkxL4qd5U88jX5oFVqAOGwHPG/owC3bLJNZmEHQpgdWQIDAQAB
AoGBAIIMQiYq86DOsbr3EVj2LQ7JzFy0u+6a40o08/gov6vKtpylpaY5z+20O0oM
b8Jwz4p8zG35ozuhWYvqoRPBqi6QRiK0cRxvphAWAVxxkeeetW4BNZvpLd8Rxaq4
S9YYOAlrV0CZihnE2GL95UU39uJAEvCvHUsxEUdK3N8XpiSBAkEA5akWLd1jWJUA
rCIwKO4SbNmNC3IEkxbX87uFTg9AxXbdd8/+oYB6AcekOJpuQqArpXk4QyO4hlFB
V6QSZfX1MQJBAOG7tr4r1nolk7Jr+OmK3Fg+7YiwCcKqw/BRextjqzbRN4OGJwOU
wB+NX9KytdWAzlWhxtJ3ZUGgwrM5TE19QKkCQD4IbAs7b2gv5xyXp3aGx8dLBAQB
aibo1q/pCNrK1+3+a1e/gMHS2CG+8Saw3/NzHBb4JTBNZ7wwGnw3vxh3VSECQQDS
NMC/JQmMI5P/kcZwjMwWLTt7jxr3uZfPIcF3RwA4gumkQ/fuwMWMXFWAWsUzdcgv
PABvG5oiXDcTOOdSPbJJAkBse1cZj4YsRvMn4yvj5aVWDkUX32w7JtXixMXnINe2
nOytiKm1pO2p4ID467mDUzla7C+RZD2DbmUAEgdn1cc8
-----END RSA PRIVATE KEY-----)";*/

  if (argc < 2) {
    std::cout << "Please enter the hostname" << std::endl;
    exit(-1);
  }
  grpc::string hostname(argv[1]);
  grpc::string address = hostname + grpc::string(":9500");

  auto contents = [](const std::string& filename) -> std::string {
        std::ifstream fh(filename);
        std::stringstream buffer;
        buffer << fh.rdbuf();
        fh.close();
        return buffer.str();
  };

  auto ssl_options = grpc::SslCredentialsOptions();
  ssl_options.pem_cert_chain = contents("/rbctrl/apiserver/certificate/client/client-cert.pem");
  ssl_options.pem_private_key = contents("/rbctrl/apiserver/certificate/client/client-key.pem");
  ssl_options.pem_root_certs = contents("/rbctrl/apiserver/certificate/ca/ca-cert.pem");

  GreeterClient greeter(grpc::CreateChannel(
      address, grpc::SslCredentials(ssl_options)));

  ::google::protobuf::int32 i = 1;
  std::string reply = greeter.SayHello(i);
  std::cout << "Greeter received: " << reply << std::endl;

  return 0;
}
