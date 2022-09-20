package main

import (
	"github.com/langwan/langgo"
	rpc "github.com/langwan/langgo/core/grpc"
	"server/pb"
	"server/service/server"
)

const addr = "localhost:8000"

// gin.new rpc.new
// gin.new gin.use gin.post(.) gin.run(addr)
func main() {
	langgo.Run()
	cg := rpc.New()
	cg.Use(rpc.LogUnaryServerInterceptor())
	pb.RegisterServerServer(cg.Server(), server.Server{})
	cg.Run(addr)
}
