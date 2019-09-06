package main

import (
	"context"
	"io"
	"log"
	"net"
	"sort"

	calculator_pb "github.com/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, in *calculator_pb.SumRequest) (*calculator_pb.SumResponse, error) {
	log.Print("Sum function invoked\n")
	log.Printf("Received request , request: %v", in)
	sum := in.GetSum().GetFirst() + in.GetSum().GetSecond()
	return &calculator_pb.SumResponse{
		Result: sum,
	}, nil
}

func (*server) PrimeNumberDecomposition(req *calculator_pb.PRequest, stream calculator_pb.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Printf("Calculate prim numbers invoked\n")
	num := req.GetNumber()
	divisor := int32(2)
	for num > 1 {
		if num%divisor == 0 {
			stream.Send(&calculator_pb.PResponse{
				Number: divisor,
			})
			num = num / divisor
		} else {
			divisor++
		}
	}
	return nil

}

func (*server) Average(stream calculator_pb.CalculatorService_AverageServer) error {
	sum := 0.0
	count := 0.0
	for {
		val, err := stream.Recv()
		if err == io.EOF {
			average := sum / count
			return stream.SendAndClose(&calculator_pb.AverageResponse{
				Number: average,
			})
		}
		if err != nil {
			log.Fatalf("Error occured , %v", err)
		}
		sort.Sort(nil)
		sum += float64(val.GetNumber())
		count++
	}
}

func main() {
	log.Print("Running calculator server\n")

	lis, err := net.Listen("tcp", "0.0.0.0:8085")
	if err != nil {
		log.Fatalf("unable to start server, error: %v", err)
	}

	s := grpc.NewServer()
	calculator_pb.RegisterCalculatorServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve. error: %v", err)
	}
}
