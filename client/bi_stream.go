package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Pawan109/grpc-mini-project/proto"
)

func callSayHelloBidirectionalStream(client pb.GreetserviceClient, names *pb.NamesList) {
	log.Printf("bidirectional streaming started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	//ab bidirectional mein client jo req bhejta hai , ab usko streams k req bhi bhejne padenge , aur streams ke response accept bhi krne padege
	//uske liye ->we make channel -> jo recv wala part dekhega

	waitc := make(chan struct{}) //channel ka datatype is struct{} here , so only struct{} type ka data hi channel k andar jaega

	go func() { //go routine
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc) //sb kuch receive krne ke baad closing the channel
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil { //ab send bhi krdiya hai
			log.Fatalf("error while sending %v", err)
		}
		time.Sleep(2 * time.Second)

	}
	stream.CloseSend() //send krna bnd krdiya hai ab client ne (req ko) , ab jo server se receive hua hoga(response) channel mein uske liye->
	<-waitc            //jo channel se hai bahar nikalo (send)
	log.Printf("Bidirectional Streaming finished")

}
