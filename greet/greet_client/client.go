package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I am a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error occured while connecting to sever: %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	// doUnary(c)
	doStream(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do a unary")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sidharth",
			LastName:  "Kamboj",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC : %v", err)
	}
	fmt.Printf("Response from greet : %v", res.Result)
	fmt.Printf("created client: %f", c)
}

func doStream(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a client streaming")

	stream, err := c.LongGreet(context.Background())

	req := []*greetpb.LGRequest{
		&greetpb.LGRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sidharth",
			},
		},
		&greetpb.LGRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sidharth2",
			},
		},
		&greetpb.LGRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sidharth3",
			},
		},
		&greetpb.LGRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sidharth4",
			},
		},
	}
	if err != nil {
		log.Fatalf("error while calling longGreet : %v", err)
	}

	for _, r := range req {
		fmt.Printf("Sending request: %v", r)
		stream.Send(r)
		time.Sleep(100 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error occured, %v", err)
	}
	fmt.Printf("response : %v", res)

}
