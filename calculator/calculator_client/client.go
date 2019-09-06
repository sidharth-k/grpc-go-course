package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	calculator_pb "github.com/grpc-go-course/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting client")
	conn, err := grpc.Dial("localhost:8085", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to server , error: %v", err)
	}
	defer conn.Close()
	c := calculator_pb.NewCalculatorServiceClient(conn)
	getAverage(c)
	// getSum(c)
	// getFactors(c)
}

func getAverage(c calculator_pb.CalculatorServiceClient) {
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error occured while calling average function, %v ", err)
	}
	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UnixNano())
		num := int32(rand.Intn(10))
		log.Printf("sending num , %v", num)
		stream.Send(&calculator_pb.AverageRequest{
			Number: num,
		})
	}
	result, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error occurred while receiving response, %v", err)
	}
	log.Printf("Received average : %v", result)
}

func getSum(c calculator_pb.CalculatorServiceClient) {
	log.Println("Starting call to server")
	req := &calculator_pb.SumRequest{
		Sum: &calculator_pb.Sum{
			First:  10,
			Second: 20,
		},
	}
	resp, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Printf("Error while getting sum from server, err:%v", err)
		return
	}
	log.Printf("Received response , response: %v", resp)
}

func getFactors(c calculator_pb.CalculatorServiceClient) {
	log.Println("Starting call to server")
	req := &calculator_pb.PRequest{
		Number: 567238,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while getting sum from server, err:%v", err)
		return
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ERROR OCCURED")
		}
		fmt.Println(res.GetNumber())
	}
	_, _ = context.WithTimeout(context.Background(), time.Second)
}
