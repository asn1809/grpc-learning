package main

import (
	"context"
	"log"

	pb "github.com/asn1809/grpc-learning/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("Starting to do a Sum RPC...")

	req := &pb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v\n", err)
	}

	log.Printf("Response from Sum: %v\n", res.Result)
}
