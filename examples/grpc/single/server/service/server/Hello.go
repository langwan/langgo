package server

import (
	"context"
	"server/pb"
)

func (s Server) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Msg: "hello"}, nil
}
