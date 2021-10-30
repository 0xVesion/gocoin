package main

import (
	"log"
	"net"

	"github.com/0xVesion/gocoin/node"
	"github.com/0xVesion/gocoin/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterNodeServer(grpcServer, node.NewServer())
	grpcServer.Serve(lis)
}
