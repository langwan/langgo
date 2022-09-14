package server

import (
	"context"
	"grpc_server/rpc"
)

func (s Server) Hello(ctx context.Context, empty *rpc.Empty) (*rpc.HelloResponse, error) {
	//TODO implement me
	panic("implement me")
}
