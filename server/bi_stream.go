package main

import (
	"io"
	"log"

	pb "github.com/Pawan109/grpc-mini-project/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.Greetservice_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv() //stream recv krne ke liye
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("got request with name : %v", req.Name)

		res := &pb.HelloResponse{ //stream mil gyi , ab usske aake "hello " likh ke send karega
			Message: "hello" + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err

		}
	}
}
