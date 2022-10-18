---
title: "GRPC"
---

# GRPC

## 简单服务器

完整例子参考 examples/grpc/single

server 端

```go
package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core/rpc"
	"server/pb"
	"server/service/server"
)

const addr = "localhost:8000"

func main() {
	langgo.Run()
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
```

client 端


```go
package main

import (
	"client/pb"
	"context"
	"fmt"
	"github.com/langwan/langgo/core/rpc"
	"google.golang.org/grpc"
	"log"
)

func main() {

	conn, err := rpc.NewClient(nil, "127.0.0.1:8000", grpc.WithInsecure())

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	ServerClient := pb.NewServerClient(conn)

	helloResponse, err := ServerClient.Hello(context.Background(), &pb.HelloRequest{From: "client"})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	log.Println(helloResponse, err)
}
```

## tls 双向认证

完整例子参考 examples/grpc/tls

server 端

```
package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core/rpc"
	"server/pb"
	"server/service/server"
)

const addr = "localhost:8000"

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
```

client 端

```
package main

import (
	"client/pb"
	"context"
	"fmt"
	"github.com/langwan/langgo/core/rpc"
	"log"
)

const addr = "localhost:8000"

func main() {

	conn, err := rpc.NewClient(&rpc.Tls{
		Crt:   "../keys/client.crt",
		Key:   "../keys/client.key",
		CACrt: "../keys/ca.crt",
	}, addr)
	if err != nil {
		panic(err)
	}

	ServerClient := pb.NewServerClient(conn)

	helloRespone, err := ServerClient.Hello(context.Background(), &pb.Empty{})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	log.Println(helloRespone, err)
}
```

## etcd 服务发现与负载均衡

完整例子参考 examples/grpc/etcd

server 端

```go
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
```

client 端

```go
package main

import (
	"client/pb"
	"context"
	"fmt"
	"github.com/langwan/langgo/core/rpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const etcdHost = "http://localhost:2379"
const serviceName = "langgo/server"

func main() {

	etcdClient, err := clientv3.NewFromURL(etcdHost)
	if err != nil {
		panic(err)
	}
	etcdResolver, err := resolver.NewBuilder(etcdClient)

	conn, err := rpc.NewClient(nil, fmt.Sprintf("etcd:///%s", serviceName), grpc.WithResolvers(etcdResolver), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
	if err != nil {
		panic(err)

	}

	ServerClient := pb.NewServerClient(conn)

	for {
		helloRespone, err := ServerClient.Hello(context.Background(), &pb.Empty{})
		if err != nil {
			fmt.Printf("err: %v", err)
			return
		}

		log.Println(helloRespone, err)
		time.Sleep(500 * time.Millisecond)
	}

}
```