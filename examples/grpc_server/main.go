package main

import (
	"github.com/langwan/langgo"
	lgrpc "github.com/langwan/langgo/components/grpc"
	"github.com/langwan/langgo/components/jwt"
	"grpc_server/rpc"
	"grpc_server/service/server"
)

func main() {
	langgo.Run(&jwt.Instance{}, &lgrpc.Instance{})
	lgrpc.New()
	rpc.RegisterServerServer(lgrpc.GetInstance().GetServer(), &server.Server{})
}
