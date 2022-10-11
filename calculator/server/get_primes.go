package main

import (
	"log"

	pb "github.com/dotNate/go-grpc-practice/calculator/proto"
)

func (s *Server) GetPrimes(in *pb.GetPrimesRequest, stream pb.CalculatorService_GetPrimesServer) error {
	log.Printf("GetPrimes was invoked with %v\n", in)

	number := in.N
	divisor := int32(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.GetPrimesResponse{
				Result: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}
