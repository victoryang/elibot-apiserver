package v2

import (
	"crypto/tls"
	"elibot-apiserver/config"
	"elibot-apiserver/auth"

	Log "elibot-apiserver/log"
	pb "elibot-apiserver/serverpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func LoadSelfSignedCert() (*credentials.ServerOption, error){
	 // Load the certificates from disk
    certificate, err := tls.LoadX509KeyPair(auth.GetCert(), auth.GetKey())
    if err != nil {
        return nil, fmt.Errorf("could not load server key pair: %s", err)
    }

    // Create a certificate pool from the certificate authority
    certPool := x509.NewCertPool()
    ca, err := ioutil.ReadFile("/rbctrl/apiserver/certificate/ca/ca-cert.pem")
    if err != nil {
        return nil, fmt.Errorf("could not read ca certificate: %s", err)
    }

    // Append the client certificates from the CA
    if ok := certPool.AppendCertsFromPEM(ca); !ok {
        return errors.New("failed to append client certs")
    }

    // Create the TLS credentials
    creds := credentials.NewTLS(&tls.Config{
    	ClientAuth:   tls.RequireAndVerifyClientCert,
    	Certificates: []tls.Certificate{certificate},
    	ClientCAs:    certPool,
    })
    return creds, nil
}

// set grpc options for optimization
func setGrpcOptions(c *config.GrpcEntryPoint) []grpc.ServerOption {
	var opts []grpc.ServerOption
	/*opts = append(opts, grpc.CustomCodec(&codec{}))*/
	if auth.IsSSL() {
		Log.Debug("grpc: ssl enabled")
		/*if cred, err := credentials.NewServerTLSFromFile(auth.GetCert(), auth.GetKey()); err==nil {
			opts = append(opts, grpc.Creds(cred))
		}*/
		if cred, err := LoadSelfSignedCert(); err==nil {
			opts = append(opts, grpc.Creds(cred))
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