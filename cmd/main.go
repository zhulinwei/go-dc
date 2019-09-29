package main

import (
	"github.com/zhulinwei/gin-demo/pkg/config"
	"github.com/zhulinwei/gin-demo/pkg/router"
	"github.com/zhulinwei/gin-demo/pkg/rpc"
	"log"
)

func main() {
	httpPort := config.ServerConfig().HttpPort
	grpcPort := config.ServerConfig().GrpcPort

	go rpc.GRPCRun(grpcPort)
	if err := router.BuildRoute().Run(httpPort); err != nil {
		log.Fatalf("server run failed: %v", err)
	}
}
