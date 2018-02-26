// Package main is gRPC client that connects to server
package main

import (
	"log"

	"github.com/giancarlopetrini/grpcdemo/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	// create connection, dial, and defer its close
	var conn *grpc.ClientConn

	// Create client TLS credentals
	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	conn, err = grpc.Dial("localhost:7777", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("could not dial: %s", err)
	}
	defer conn.Close()

	// create grpc client for the registered ping service
	c := api.NewPingClient(conn)

	// send PingMessage, setting foo to greeting as defined in proto
	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "foo"})
	if err != nil {
		log.Fatalf("error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)
}
