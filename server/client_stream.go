package main

import (
	"io"
	"log"

	pb "github.com/Pawan109/grpc-mini-project/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.Greetservice_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("got request with name: %v", req.Name)
		messages = append(messages, "Hello "+req.Name)
	}
}
