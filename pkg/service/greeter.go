package service

import (
	"context"
	"github.com/zhulinwei/go-dc/pkg/rpc"
	greeterPb "github.com/zhulinwei/grpc-demo/helloworld/greeter/proto"
	"google.golang.org/grpc"
	"time"
)

type IGreeter interface {
	QueryGreeterFromGrpc(name string) (*greeterPb.HelloReply, error)
}

type GreeterClient struct {
	Client *grpc.ClientConn
}

func BuildGreeterService() IGreeter {
	return GreeterClient{
		Client: rpc.BuildGrpc().Client1(),
	}
}

func (greeter GreeterClient) QueryGreeterFromGrpc(name string) (*greeterPb.HelloReply, error) {
	greeterClient := greeterPb.NewGreeterClient(greeter.Client)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return greeterClient.SayHello(ctx, &greeterPb.HelloRequest{Name: name})
}
