package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Pawan109/grpc-mini-project/proto"
)

func callSayHelloServerStream(client pb.GreetserviceClient, names *pb.NamesList) {
	log.Printf("streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("couldn't send namesL %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("err while streaming %v", err)
		}
		log.Println(message)
	}
	log.Printf("streaming finished ")
}
