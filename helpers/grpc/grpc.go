package helperGrpc

import (
	"google.golang.org/grpc"
)

// from https://github.com/grpc-ecosystem/go-grpc-middleware/blob/master/chain.go
//
// ChainUnaryServer creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainUnaryServer(one, two, three) will execute one before two before three, and three
// will see context changes of one and two.
//
// While this can be useful in some scenarios, it is generally advisable to use google.golang.org/grpc.ChainUnaryInterceptor directly.

// ChainStreamServer creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainUnaryServer(one, two, three) will execute one before two before three.
// If you want to pass context between interceptors, use WrapServerStream.
//
// While this can be useful in some scenarios, it is generally advisable to use google.golang.org/grpc.ChainStreamInterceptor directly.
func ChainStreamServer(interceptors ...grpc.StreamServerInterceptor) grpc.StreamServerInterceptor {
	n := len(interceptors)

	// Dummy interceptor maintained for backward compatibility to avoid returning nil.
	if n == 0 {
		return func(srv interface{}, stream grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			return handler(srv, stream)
		}
	}

	if n == 1 {
		return interceptors[0]
	}

	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		currHandler := handler
		for i := n - 1; i > 0; i-- {
			innerHandler, i := currHandler, i
			currHandler = func(currentSrv interface{}, currentStream grpc.ServerStream) error {
				return interceptors[i](currentSrv, currentStream, info, innerHandler)
			}
		}
		return interceptors[0](srv, stream, info, currHandler)
	}
}
