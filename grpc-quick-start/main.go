package main

import (
	"log"
	"net"

	server "github.com/dcastrobianca/grpc/grpc/server"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := server.CreateServer()

	log.Println("gRPC server running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
