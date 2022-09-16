package main

import (
	"github.com/langwan/langgo"
	rpc "github.com/langwan/langgo/core/grpc"
	"google.golang.org/grpc"
	"server/pb"
	"server/service/server"
)

const addr = "localhost:8000"

func main() {
	langgo.Run()
	cg := rpc.New(grpc.MaxSendMsgSize(1024*1024*10), grpc.MaxRecvMsgSize(1024*1024*10))
	cg.Use(rpc.LogUnaryServerInterceptor())
	pb.RegisterServerServer(cg.Server(), server.Server{})
	cg.Run(addr)
}
