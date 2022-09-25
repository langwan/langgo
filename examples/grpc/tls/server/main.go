package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core/rpc"
	"server/pb"
	"server/service/server"
)

const addr = "localhost:8000"

// gin.new rpc.new
// gin.new gin.use gin.post(.) gin.run(addr)
func main() {
	langgo.Run()
	cg := rpc.NewServer(&rpc.Tls{
		Crt:   "../keys/server.crt",
		Key:   "../keys/server.key",
		CACrt: "../keys/ca.crt",
	})
	cg.Use(rpc.LogUnaryServerInterceptor())
	gs, err := cg.Server()
	if err != nil {
		panic(err)
	}
	pb.RegisterServerServer(gs, server.Server{})
	cg.Run(addr)
}
