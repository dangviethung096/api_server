package main

import (
	"api_server_core/example/server"
	"fmt"
	"log"
	"net"
	"server_proto"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:11001")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	server_proto.RegisterExampleServiceServer(grpcServer, server.NewServer())

	fmt.Println("Start example server in port: 11001")
	grpcServer.Serve(lis)
}
