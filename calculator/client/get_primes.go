package main

import (
	"context"
	"io"
	"log"

	pb "github.com/dotNate/go-grpc-practice/calculator/proto"
)

func doGetPrimes(c pb.CalculatorServiceClient) {
	log.Println("doGetPrimes was invoked")

	req := &pb.GetPrimesRequest{
		N: 120,
	}

	stream, err := c.GetPrimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GetPrimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GetPrimes: %d\n", msg.Result)
	}
}
