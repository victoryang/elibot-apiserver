package v2

import (
	pb "elibot-apiserver/serverpb"
	"google.golang.org/grpc"
)

// TODO:
/*func setGrpcOptions() []grpc.ServerOption {
	var opts []grpc.ServerOption
	opts = append(opts, grpc.CustomCodec(&codec{}))
	if tls != nil {
		opts = append(opts, grpc.Creds(credentials.NewTLS(tls)))
	}
	opts = append(opts, grpc.UnaryInterceptor(newUnaryInterceptor(s)))
	opts = append(opts, grpc.StreamInterceptor(newStreamInterceptor(s)))
	opts = append(opts, grpc.MaxRecvMsgSize(int(s.Cfg.MaxRequestBytes+grpcOverheadBytes)))
	opts = append(opts, grpc.MaxSendMsgSize(maxSendBytes))
	opts = append(opts, grpc.MaxConcurrentStreams(maxStreams))
	return opts
}*/

func Server() *grpc.Server {
	/*opts := setGrpcOptions()*/
	s := grpc.NewServer()

	/* An hello example*/
	pb.RegisterGreeterServer(s, NewHelloServer())

	/*TODO: Register all servers here*/

	return s
}