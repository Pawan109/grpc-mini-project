package main

import (
	"log"
	"time"

	pb "github.com/Pawan109/grpc-mini-project/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.Greetservice_SayHelloServerStreamingServer) error {
	log.Printf("got request with names :%v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "hello " + name,
		}
		if err := stream.Send(res); err != nil { //serverstream mein -> server send krra hai , //clientstream mein -> client send krr hai
			return err
		}
		// 2 second delay to simulate a long running process
		time.Sleep(2 * time.Second)
	}
	return nil
}
