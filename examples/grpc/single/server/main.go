package main

import (
	"github.com/langwan/langgo"
	rpc "github.com/langwan/langgo/core/grpc"
	"server/pb"
	"server/service/server"
)

const addr = "localhost:8000"

func main() {
	langgo.Run()
	cg := rpc.New()
	cg.Use(rpc.LogUnaryServerInterceptor())
	pb.RegisterServerServer(cg.Server(), server.Server{})
	cg.Run(addr)
}
