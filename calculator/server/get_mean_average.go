package main

import (
	"io"
	"log"

	pb "github.com/dotNate/go-grpc-practice/calculator/proto"
)

func (s *Server) GetMeanAverage(stream pb.CalculatorService_GetMeanAverageServer) error {
	log.Println("GetMeanAverage was invoked")

	res := float32(0.0)
	iterator := float32(0.0)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GetMeanAverageResponse{
				Result: float32(res) / float32(iterator),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)

		iterator++
		res += float32(req.Num)
	}
}
