package main

import (
	"log"

	pb "github.com/Pawan109/grpc-mini-project/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	//to connect to the server -> grpc.Dial
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetserviceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Pawan", "Alice", "Bob"},
	}

	//callSayHello(client) //done
	//callSayHelloServerStream(client, names) //client making request -> server should send the streams now
	//callSayHelloClientStream(client, names)
	callSayHelloBidirectionalStream(client, names)
}
