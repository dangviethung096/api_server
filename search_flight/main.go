package main

import (
	"fmt"
	"log"
	"net"
	"search_flight/server"

	pb "server_proto"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:11002")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSearchFlightServiceServer(grpcServer, server.NewServer())

	fmt.Println("Start server at port: 11002!")
	grpcServer.Serve(listener)
}
