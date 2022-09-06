package helperGrpc

import (
	"context"
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
