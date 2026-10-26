package main

import (
	"io"
	"log"

	pb "github.com/asn1809/grpc-learning/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone function was invoked\n")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		res := "Hello " + req.FirstName + "!"
		sendErr := stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}
}
