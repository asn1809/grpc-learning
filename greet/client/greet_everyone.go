package main

import (
	"context"
	pb "github.com/asn1809/grpc-learning/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(client pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GreetEveryone: %v", err)
	}

	requests := []*pb.GreetRequest{
		{FirstName: "John"},
		{FirstName: "Jane"},
		{FirstName: "Alice"},
		{FirstName: "Bob"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			log.Printf("Sending message: %v", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			log.Printf("Received: %v", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc
}
