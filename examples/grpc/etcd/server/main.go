package main

import (
	"flag"
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/core/rpc"
	"os"
	cs "server/components/server"
	"server/pb"
	"server/service/server"
	"syscall"
)

func main() {

	var port int
	flag.IntVar(&port, "port", 8001, "port")
	flag.Parse()
	addr := fmt.Sprintf("localhost:%d", port)

	langgo.Run(&cs.Instance{})
	core.SignalHandle(&core.SignalHandler{
		Sig: syscall.SIGINT,
		F: func() {
			rpc.EtcdUnRegister(cs.GetInstance().ServiceName, addr)
			os.Exit(int(syscall.SIGINT))
		},
	})
	defer func() {
		core.DeferRun()
	}()
	rpc.EtcdRegister(cs.GetInstance().EtcdHost, cs.GetInstance().ServiceName, addr, 50)
	cg := rpc.NewServer(nil)
	cg.Use(rpc.LogUnaryServerInterceptor())
	gs, err := cg.Server()
	if err != nil {
		panic(err)
	}
	pb.RegisterServerServer(gs, server.Server{})
	err = cg.Run(addr)
	if err != nil {
		panic(err)
	}
}
