package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"vesion.io/go/gocoin/node"
	"vesion.io/go/gocoin/pb"
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
