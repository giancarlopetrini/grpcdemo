package main

import (
	"fmt"
	"log"
	"net"

	"github.com/giancarlopetrini/grpcdemo/api"
	"google.golang.org/grpc"
)

// main starts a gRPC server and waits for connection
func main() {
	// create listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprint("localhost:", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server instance using struct from api/handler
	s := api.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Ping Service to the server
	api.RegisterPingServer(grpcServer, &s)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %s", err)
	}
}
