package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dotNate/go-grpc-practice/calculator/proto"
)

func doGetMeanAverage(c pb.CalculatorServiceClient) {
	log.Println("goGetMeanAverage was invoked")

	reqs := []*pb.GetMeanAverageRequest{
		{Num: int32(1)},
		{Num: int32(2)},
		{Num: int32(3)},
		{Num: int32(4)},
	}

	stream, err := c.GetMeanAverage(context.Background())

	if err != nil {
		log.Fatalf("Error while calling GetMeanAverage: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from GetMeanAverage: %v\n", err)
	}

	log.Printf("GetMeanAverage: %f\n", res.Result)
}
