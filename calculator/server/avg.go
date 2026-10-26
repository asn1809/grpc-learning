package main

import (
	"io"
	"log"

	pb "github.com/asn1809/grpc-learning/calculator/proto"
)

func (s *Server) Avg(server pb.CalculatorService_AvgServer) error {
	log.Printf("Avg function was invoked")

	var sum int32 = 0
	count := 0

	for {
		req, err := server.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return server.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Println("Receiving number: %v\n", req.Number)
		sum += req.Number
		count++
	}
}
