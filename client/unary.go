package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Pawan109/grpc-mini-project/proto"
)

func callSayHello(client pb.GreetserviceClient) { //sirf greetService hi lega as a parameter
	//ctx necessary for req,res -> time out laga diya
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() //aur pura fn khtm hone ke baad last mein call kiya

	res, err := client.SayHello(ctx, &pb.NoParam{}) //SayHelo func ko Server ke unary mein define kiya hai
	//usme pass ctx & NoParam (kyuki mko kuch ni bhejna , bss ek req bhejni hai, jiska heloo wala response aaega )

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("%s", res.Message)
}
