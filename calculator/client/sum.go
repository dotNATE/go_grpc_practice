package main

import (
	"context"
	"log"

	pb "github.com/dotNate/go-grpc-practice/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		N1: 5,
		N2: 8,
	})

	if err != nil {
		log.Fatalf("Could not get sum: %v\n", err)
	}

	log.Printf("Sum: %v\n", res.Result)
}
