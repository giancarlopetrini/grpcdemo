// api/hadler.go

package api

import (
	"log"

	"golang.org/x/net/context"
)

// Server type (exportable) will be gRCP server
type Server struct {
}

// SayHello generates responses to Ping quest
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}
