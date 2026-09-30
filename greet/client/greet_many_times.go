package main

import (
	"context"
	pb "github.com/asn1809/grpc-learning/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("Starting to do a GreetManyTimes RPC...")

	req := &pb.GreetRequest{
		FirstName: "Narasago",
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v\n", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while receiving stream: %v\n", err)
		}

		log.Printf("Response from GreetManyTimes: %v\n", msg.GetResult())
	}
}
