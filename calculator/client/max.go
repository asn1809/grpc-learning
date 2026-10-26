package main

import (
	"context"
	pb "github.com/asn1809/grpc-learning/calculator/proto"
	"io"
	"log"
	"time"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Max: %v", err)
	}

	requests := []int32{1, 5, 3, 6, 2, 20}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			log.Printf("Sending message: %v", req)
			stream.Send(&pb.MaxRequest{
				Number: req,
			})
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
			log.Printf("Received a new maximum: %v", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc
}
