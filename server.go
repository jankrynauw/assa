package main

import (
	pb "github.com/jankrynauw/assa/protobufs"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	log.Printf("starting server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen: %v", err)
	}

	// Instantiate gRPC server object
	grpcServer := grpc.NewServer()
	pb.RegisterMembersServer(grpcServer, &myService{})

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
