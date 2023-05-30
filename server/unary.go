package main

import (
	"context"

	pb "github.com/Pawan109/grpc-mini-project/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	//do cheezein lega -> context (needed for client-server / req-res  & a req with no parameters) & will return a hello response& err
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
