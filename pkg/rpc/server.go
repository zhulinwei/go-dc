package rpc

import (
	"github.com/zhulinwei/go-dc/pkg/rpc/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
)

func GRPCRun(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	protobuf.RegisterGreeterServer(server, &GreeterServer{})
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
