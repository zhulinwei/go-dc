package rpc

import (
	"context"
	"github.com/zhulinwei/gin-demo/pkg/model/protobuf"
)

type GreeterServer struct{}

func (g *GreeterServer) SayHello(ctx context.Context, req *protobuf.HelloRequest) (*protobuf.HelloReply, error) {
	return &protobuf.HelloReply{Message: "Hello " + req.Name}, nil
}