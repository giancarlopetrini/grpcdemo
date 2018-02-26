// Package api handles ping request
package api

import (
	"log"

	"golang.org/x/net/context"
)

// Server represents gRPC server
type Server struct{}

// SayHello - generates response to Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}
