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
