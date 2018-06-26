#include <iostream>
#include <memory>
#include <string>

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
  grpc::string cert = R"(-----BEGIN CERTIFICATE-----
MIIDhTCCAm2gAwIBAgIJAOLZRxLasLlqMA0GCSqGSIb3DQEBCwUAMFkxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQxEjAQBgNVBAMMCWVsaWJvdC5jbjAeFw0xODA2MjYwNjMy
MjBaFw0xOTA2MjYwNjMyMjBaMFkxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21l
LVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQxEjAQBgNV
BAMMCWVsaWJvdC5jbjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAOH9
QOViCmFRTZycaHVhz/WObG/oFms5L75r0vaKMoT1PC7TPLntDLjClWhp55gqfkbg
7lBe0hPO8kd3rV5Os88H6RgZtEyc5MEQaK++BhnoZR5B3Wpe4nbvcn0Q5BLF1s3C
L5wj3gnmAqe97v8jyHH5ij1V/+bcer1B9Z5XG2fj0PBkLWI2YPdjx8Zjf0sln8az
l6tGgrLlUy6hFl64DU18TJxRmifizVNEraYl07CTbylS28QzCNBGPnlnciQRx3Ai
k4bCX3+uhJ4mBqQEwJrzuRP/ZS3ZsBXm2fLWujYqld/RgPIhn+Bb3xY9JABDAdMZ
ZBi5f63ysla9+cKXPscCAwEAAaNQME4wHQYDVR0OBBYEFDM+qxxx8+Qc3UJ2eAyw
JB+atkMxMB8GA1UdIwQYMBaAFDM+qxxx8+Qc3UJ2eAywJB+atkMxMAwGA1UdEwQF
MAMBAf8wDQYJKoZIhvcNAQELBQADggEBAIgrvpoxqKGGT1WwjfzTTY7qw71Czbxj
zXNPDEmu8G78wKhN0yHy9vZDUuRO/bK+O2yZqy0VyGkbSOp+q2NgJoI38zc65YXT
9+WrwnQUg/eF8xXtYSeSbRGaZCCw7LGejheVqlawxJ/OSUixBLBxiPxJDtxcXZVX
F8yaPyaoBVNGiGid7w9HQmOwLFw94e+j3cX5RizbneIlyspI2JAdeCWPFw5DsCRJ
AQw/yAlfzAp1TwxRiXY8nlgzc/zZ9+mpm++Q76/tRcVECw7wtZp+hJQe6MiWT8kG
urAazDpZHUmhTxyRMb8b77Z4ERnhFtj8BqsG89lkxEDzkbzKGRhkxlo=
-----END CERTIFICATE-----)";

  grpc::string key = R"(-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA4f1A5WIKYVFNnJxodWHP9Y5sb+gWazkvvmvS9ooyhPU8LtM8
ue0MuMKVaGnnmCp+RuDuUF7SE87yR3etXk6zzwfpGBm0TJzkwRBor74GGehlHkHd
al7idu9yfRDkEsXWzcIvnCPeCeYCp73u/yPIcfmKPVX/5tx6vUH1nlcbZ+PQ8GQt
YjZg92PHxmN/SyWfxrOXq0aCsuVTLqEWXrgNTXxMnFGaJ+LNU0StpiXTsJNvKVLb
xDMI0EY+eWdyJBHHcCKThsJff66EniYGpATAmvO5E/9lLdmwFebZ8ta6NiqV39GA
8iGf4FvfFj0kAEMB0xlkGLl/rfKyVr35wpc+xwIDAQABAoIBAHjytL/PZ+TLwbD2
2MUx8z3l2s1NtmnuckoEpSTDs+QK6Anbjh2n6+4aXiJCw+thmrHOgl1LUj29vVdY
itX1YzCDeFIot6FlOtzFLnIYTmWhRjUaZSkf3S5wWqLq3lXZwmve6OI1JAnCK1uh
QH8WpJmzRzQqicoB6ELQQWdcEVPYzShA8c31WC5Q+C/J1qH+DySE7h2PTp8BB/LL
g9J1hfblY/aCNq6saWlFzRo4mPFPyHHvzKwzzdW3+cc1wklo0HGqAIbEM1L+YFb/
6rQe8XbA1qihj4mcA6kuWP1dqiruv7FtT4b6r+9oAE1utQMO/yq0Ud31ns6s0gXc
/PjMC3ECgYEA84HAKROca/HnCqCnQMu1B+/cxgvA1R5PvmAr9dInXoXa5nweX1kF
dfsP5XArTZfO2+aAfFckzyY3CMzJSMxA9BVEvTznaGcnmdnoFeD99LVw0ys+tZXQ
aicd23Jtdzx0gRy57Trt5t7mfXhsWWbigjG9Aj7kYY6DjdtT/fbl0pkCgYEA7ZVs
yZyuZG4xSgF32k2g+ZBZKSUZUi87BSBpH6iplSa1E71Jnt2zuyJKHt4VTSVlgxx1
Uz6af2YneVpwPfmeYOEmrRi8yN8/2uHI9b8cGNG1OCAAd7E8AupIJIH7n/D0F9dA
PPRbF5pOpkdqBfI1bXdWyYxLoKKQeGd7N4Wl2F8CgYB6mywMAqnR2B0Cxt0vzOG7
u1QoC6buJ7LzIi1AOq5D85XCU2Bflc0gGGdqmM3U5cjsA+VRtfb8rjsrnSEEHIPf
4g2YUuAZO0c/Oe6XhY+Y9Pp3+OA8QdCMgmGQKs9fJ/tpPvOGtRMwGa9oIYg8g4ct
EqoeRVhsnnsyo+pohzY0kQKBgQDmTPpvVNoXsFlHC7VKgAWS5UIIiFXLPM1RYuGE
NtJsKmFNCSfcP7yBofOHiG/NoHqOZX+1efH5nOSW1fwHl3jXIGmFUX1umjojysoq
rV5nEKFCDoNlgwBRMxlLilSH7eIvWhCDBbtnXcTvmjpEGU2BFzWhmcWVqP+yN11R
rNPesQKBgGDWspWqswvZhDjWG353vRx9RMUrys3WsvN0WxabenoJgsZ5RQHzh3Br
XR4NXXi7OBw7S8SXxQ/25QkJFP6FL/1YW75z5TM2agePYXVQnaWSWFxd0WDwOjWK
jmCry0kFqJRGaq5NjYXrqqtQxJfWKLmZG6gev43iTwSdDzxvCcfh
-----END RSA PRIVATE KEY-----)";

  GreeterClient greeter(grpc::CreateChannel(
      "localhost:9500", grpc::SslCredentials(grpc::SslCredentialsOptions{"", key, cert})));

  ::google::protobuf::int32 i = 1;
  std::string reply = greeter.SayHello(i);
  std::cout << "Greeter received: " << reply << std::endl;

  return 0;
}
