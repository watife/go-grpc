package main

import (
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

	fmt.Printf("Created client: %f", c)
}
