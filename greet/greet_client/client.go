package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fakorede-bolu/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I am a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connet: %v", err)
	}

	defer conn.Close() //defer to make this execute at the very end of the func

	c := greetpb.NewGreetClient(conn)

	doUnary(c)
}

func doUnary(c greetpb.GreetClient) {
	fmt.Println("Starting to do a Unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Boluwatife",
			LastName:  "Fakorede",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}
