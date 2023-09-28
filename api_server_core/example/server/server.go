package server

import (
	"context"
	"server_proto"
	"strings"
)

func NewServer() *Server {
	return &Server{}
}

type Server struct {
	server_proto.UnimplementedExampleServiceServer
}

func (s *Server) SayHello(ctx context.Context, msg *server_proto.ExampleMessage) (*server_proto.ExampleMessage, error) {
	retValue := &server_proto.ExampleMessage{
		Message: "So impolite",
	}
	if strings.ToLower(msg.Message) == "hi" {
		retValue.Message = "Hello there! this is a welcome from server!"
	}
	return retValue, nil
}
