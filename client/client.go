package main

import (
	"context"
	"fmt"
	"log"

	"time/time"

	"google.golang.org/grpc"
)

func main() {
	// Creat a virtual RPC Client Connection on port  8008 WithInsecure (because  of http)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8008", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	// Defer means: When this function returns, call this method (meaing, one main is done, close connection)
	defer conn.Close()

	//  Create new Client from generated gRPC code from proto
	c := time.NewGetCurrentTimeClient(conn)
	SendGetTimeRequest(c)
}

func SendGetTimeRequest(c time.GetCurrentTimeClient) {
	// Between the curly brackets are nothing, because the .proto file expects no input.
	message := time.GetTimeRequest{}

	response, err := c.GetTime(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetCourse: %s", err)
	}

	fmt.Printf("Response from the Server: %s\n", response.Reply)
}
