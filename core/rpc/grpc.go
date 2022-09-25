package rpc

import (
	"context"
	ctls "crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/langwan/langgo/core/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"net"
	"runtime/debug"
	"time"
)

type Server struct {
	//server *grpc.Server
	opt        []grpc.ServerOption
	middleware []grpc.UnaryServerInterceptor
	server     *grpc.Server
	tls        *Tls
}

type Tls struct {
	Crt   string
	Key   string
	CACrt string
}

func (s *Server) Use(middleware ...grpc.UnaryServerInterceptor) {
	s.middleware = append(s.middleware, middleware...)
}

func NewClient(tls *Tls, addr string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
	if tls != nil {
		certificate, err := ctls.LoadX509KeyPair(tls.Crt, tls.Key)
		if err != nil {
			return nil, err
		}

		certPool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(tls.CACrt)
		if err != nil {
			return nil, err
		}
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			return nil, errors.New("AppendCertsFromPEM is false")
		}

		creds := credentials.NewTLS(&ctls.Config{
			Certificates:       []ctls.Certificate{certificate},
			RootCAs:            certPool,
			InsecureSkipVerify: true,
		})

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	return grpc.Dial(addr, opts...)

}

func NewServer(tls *Tls, opt ...grpc.ServerOption) *Server {
	s := &Server{
		opt: opt,
		tls: tls,
	}
	return s
}

func (s *Server) Server() (server *grpc.Server, err error) {
	if s.server == nil {
		s.opt = append(s.opt, grpc.UnaryInterceptor(ChainUnaryServer(s.middleware...)))
		if s.tls != nil {

			certificate, err := ctls.LoadX509KeyPair(s.tls.Crt, s.tls.Key)
			if err != nil {
				return nil, err
			}

			certPool := x509.NewCertPool()
			ca, err := ioutil.ReadFile(s.tls.CACrt)
			if err != nil {
				return nil, err
			}
			if ok := certPool.AppendCertsFromPEM(ca); !ok {
				panic("AppendCertsFromPEM failed")
			}

			creds := credentials.NewTLS(&ctls.Config{
				Certificates: []ctls.Certificate{certificate},
				ClientAuth:   ctls.RequireAndVerifyClientCert,
				ClientCAs:    certPool,
			})

			s.opt = append(s.opt, grpc.Creds(creds))
		}
		s.server = grpc.NewServer(s.opt...)
	}
	return s.server, nil
}

func (s *Server) Run(addr string) error {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	log.Logger("grpc", "run").Info().Str("addr", addr).Msg("server ready")
	if err := s.server.Serve(lis); err != nil {
		log.Logger("grpc", "run").Error().Err(err).Send()
		return err
	}
	return nil
}

func ChainUnaryServer(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	n := len(interceptors)

	// Dummy interceptor maintained for backward compatibility to avoid returning nil.
	if n == 0 {
		return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			return handler(ctx, req)
		}
	}

	// The degenerate case, just return the single wrapped interceptor directly.
	if n == 1 {
		return interceptors[0]
	}

	// Return a function which satisfies the interceptor interface, and which is
	// a closure over the given list of interceptors to be chained.
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		currHandler := handler
		// Iterate backwards through all interceptors except the first (outermost).
		// Wrap each one in a function which satisfies the handler interface, but
		// is also a closure over the `info` and `handler` parameters. Then pass
		// each pseudo-handler to the next outer interceptor as the handler to be called.
		for i := n - 1; i > 0; i-- {
			// Rebind to loop-local vars so they can be closed over.
			innerHandler, i := currHandler, i
			currHandler = func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
				return interceptors[i](currentCtx, currentReq, info, innerHandler)
			}
		}
		// Finally return the result of calling the outermost interceptor with the
		// outermost pseudo-handler created above as its handler.
		return interceptors[0](ctx, req, info, currHandler)
	}
}

func LogUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		st := time.Now()
		defer func() {
			if recoverError := recover(); recoverError != nil {
				log.Logger("grpc", "log").Error().Interface("recover", recoverError).Bytes("stack", debug.Stack()).Interface("req", req).Interface("resp", resp).Err(err).TimeDiff("runtime", time.Now(), st).Send()
			} else {
				log.Logger("grpc", "log").Info().Interface("req", req).Interface("resp", resp).Err(err).TimeDiff("runtime", time.Now(), st).Send()
			}
		}()
		resp, err = handler(ctx, req)
		return resp, err
	}
}
