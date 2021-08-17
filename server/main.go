package main

import (
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	pb "https://github.com/kousuketk/gRPC_chat/grpc/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9090"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
}