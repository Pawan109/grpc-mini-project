package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Pawan109/grpc-mini-project/proto"
)

func callSayHelloClientStream(client pb.GreetserviceClient, names *pb.NamesList) {
	log.Printf("client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)

	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{ //req define ki , ki req ek stream of names hai
			Name: name,
		}
		if err := stream.Send(req); err != nil { //agar req ki streams bhejne mein dikkat toh handle err
			log.Fatalf("error while sending %v", err)
		}
		log.Printf("sent request with name :%s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("client streaming finished")
	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}
	log.Printf("%v", res.Messages)
}
