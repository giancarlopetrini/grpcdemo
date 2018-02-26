// client/main.go

package main

import (
	"log"

	"github.com/giancarlopetrini/grpcdemo/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// create client connection type
	var conn *grpc.ClientConn

	// dial to server
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	// defer close of connection
	defer conn.Close()

	// init client, passing in conn, and using client generated from proto
	c := api.NewPingClient(conn)

	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "foo"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)
}
