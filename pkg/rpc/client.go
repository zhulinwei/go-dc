package rpc

import (
	"sync"

	"github.com/zhulinwei/go-dc/pkg/config"
	"github.com/zhulinwei/go-dc/pkg/util/log"
	"google.golang.org/grpc"
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
				log.Error("grpc dial fail: %v", log.String("error", err.Error()))
				return
			}
			grpcClientMap[grpcConfig.Name] = client
		}
	})
}
