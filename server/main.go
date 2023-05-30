package main

import (
	"log"
	"net" //provides a portable interface for network I/O, including TCP/IP, UDP....

	pb "github.com/Pawan109/grpc-mini-project/proto"
	"google.golang.org/grpc"
)

//define port
const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetserviceServer //ye service daali hai in helloServer
}

func main() {
	//listen on the port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start server %v", err)
	}

	//creating a grpc server
	grpcServer := grpc.NewServer()

	//register the greet Service
	pb.RegisterGreetserviceServer((grpcServer), &helloServer{})

	log.Printf("server started at %v", lis.Addr())
	//lis is the port the grpc server needs to start
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start : %v", err)
	}

}
