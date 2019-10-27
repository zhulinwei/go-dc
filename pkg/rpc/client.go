package rpc

import (
	"github.com/zhulinwei/go-dc/pkg/config"
	"google.golang.org/grpc"
	"log"
	"sync"
)

var grpcOnce sync.Once
var grpcMutex sync.Mutex
var grpcClientMap map[string]*grpc.ClientConn

type IGrpc interface {
	Client1() *grpc.ClientConn
	Client2() *grpc.ClientConn
}

type Grpc struct {
	ClientMap map[string]*grpc.ClientConn
}

func BuildGrpc() IGrpc {
	initGrpc()
	return Grpc{
		ClientMap: grpcClientMap,
	}
}

func (grpc Grpc) Client1() *grpc.ClientConn {
	return grpc.ClientMap["grpc1"]
}

func (grpc Grpc) Client2() *grpc.ClientConn {
	return grpc.ClientMap["grpc2"]
}

func initGrpc() {
	grpcConfigs := config.ServerConfig().Grpc
	grpcOnce.Do(func() {
		grpcMutex.Lock()
		defer grpcMutex.Unlock()

		grpcClientMap = make(map[string]*grpc.ClientConn, len(grpcConfigs))
		for _, grpcConfig := range grpcConfigs {
			var err error
			var client *grpc.ClientConn
			if client, err = grpc.Dial(grpcConfig.Addr, grpc.WithInsecure()); err != nil {
				log.Fatal("grpc dial fail: %v", err)
				return
			}
			grpcClientMap[grpcConfig.Name] = client
		}
	})
}
