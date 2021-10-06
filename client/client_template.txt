package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"

    "myFolder/myPackage"
)

func main() {
	// Creat a virtual RPC Client Connection on port  9080 WithInsecure (because  of http)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	// Defer means: When this function returns, call this method (meaing, one main is done, close connection)
	defer conn.Close()

	//  Create new Client from generated gRPC code from proto
	c := myPackage.NewGetXXXClient(conn)
 
	SendRequest(c)
}

func SendRequest(c myPackage.XXXClient) {
    // Between the curly brackets are nothing, because the .proto file expects no input.
	message := myPackage.somethingsomethingRequest{}

    response, err := c.somethingsomething(context.Background(), &message)
    if err != nil {
        log.Fatalf("Error when calling XXX: %s", err)
    }

    fmt.Printf("Response from the Server: %s \n", response.Reply)
}