package main

import (
	"github.com/zhulinwei/gin-demo/pkg/model/protobuf"
	"github.com/zhulinwei/gin-demo/pkg/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct {

}

func (GRPCServer) Run(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	protobuf.RegisterGreeterServer(server, &service.GreeterServer{})
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
