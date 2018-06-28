package v2

import (
	"elibot-apiserver/config"
	"elibot-apiserver/auth"

	Log "elibot-apiserver/log"
	pb "elibot-apiserver/serverpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// set grpc options for optimization
func setGrpcOptions(c *config.GrpcEntryPoint) []grpc.ServerOption {
	var opts []grpc.ServerOption
	/*opts = append(opts, grpc.CustomCodec(&codec{}))*/
	if auth.IsSSL() {
		Log.Debug("grpc: ssl enabled")
		if option, err := LoadSelfSignedCert(); err==nil {
			opts = append(opts, option)
		} else {
			Log.Debug("grpc ssl load cert fails: ", err)
			Log.Debug("grpc set unsecure")
		}
	}

	/*opts = append(opts, grpc.UnaryInterceptor(newUnaryInterceptor(s)))
	opts = append(opts, grpc.StreamInterceptor(newStreamInterceptor(s)))
	// default 4MB
	opts = append(opts, grpc.MaxRecvMsgSize(int(s.Cfg.MaxRequestBytes+grpcOverheadBytes)))
	opts = append(opts, grpc.MaxSendMsgSize(maxSendBytes))

	opts = append(opts, grpc.MaxConcurrentStreams(maxStreams))*/
	return opts
}

func Server(c *config.GrpcEntryPoint) *grpc.Server {
	opts := setGrpcOptions(c)

	s := grpc.NewServer(opts...)

	/* An hello example*/
	pb.RegisterGreeterServer(s, NewHelloServer())

	/*TODO: Register all servers here*/

	/* Register reflection service on GRPC server*/
	reflection.Register(s)
	return s
}