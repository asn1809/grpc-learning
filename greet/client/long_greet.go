package main

import (
	"context"
	pb "github.com/asn1809/grpc-learning/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("Starting to do a Long Greet RPC...")

	reqs := []*pb.GreetRequest{
		{FirstName: "John"},
		{FirstName: "Paul"},
		{FirstName: "George"},
		{FirstName: "Ringo"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet RPC: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet RPC: %v\n", err)
	}

	log.Printf("LongGreet Response: %v\n", res)
}
