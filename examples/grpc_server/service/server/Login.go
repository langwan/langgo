package server

import (
	"context"
	"grpc_server/rpc"
)

func (s Server) Login(ctx context.Context, request *rpc.SignInRequest) (*rpc.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}
