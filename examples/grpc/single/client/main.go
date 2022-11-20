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

	helloResponse, err := ServerClient.Hello(context.Background(), &pb.Empty{})

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	log.Println(helloResponse, err)
}
