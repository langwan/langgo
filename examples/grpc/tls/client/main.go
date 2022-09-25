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
