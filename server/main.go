package main

import (
	"log"
	"net"

	"github.com/giancarlopetrini/grpcdemo/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// main start a gRPC server and waits for connection
func main() {
	// create a listener on tcp port 7777
	lis, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a server instance
	s := api.Server{}

	// create TLS credentials on server
	creds, err := credentials.NewServerTLSFromFile("../cert/server.crt", "../cert/server.key")
	if err != nil {
		log.Fatalf("could not load TLS Keys: %s", err)
	}

	// create an array of gRPC options with credentials
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	// create a gRPC server object using variadic function
	grpcServer := grpc.NewServer(opts...)

	// attach ping service to the server
	api.RegisterPingServer(grpcServer, &s)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %s", err)
	}
}
