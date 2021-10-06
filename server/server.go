package main

import (
	"context"
	"fmt"
	"log"
	"net"
	t "time"
	"time/time"

	"google.golang.org/grpc"
)

type Server struct {
	time.UnimplementedGetCurrentTimeServer
}

func (s *Server) GetTime(ctx context.Context, in *time.GetTimeRequest) (*time.GetTimeReply, error) {
	fmt.Printf("Received GetTime request\n")
	return &time.GetTimeReply{Reply: t.Now().String()}, nil
}

func main() {
	// Create listener tcp on port 8008
	list, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("Failed to listen on port 8008: %v", err)
	}
	grpcServer := grpc.NewServer()
	time.RegisterGetCurrentTimeServer(grpcServer, &Server{})

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
