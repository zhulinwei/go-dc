package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/go-dc/pkg/config"
	"github.com/zhulinwei/go-dc/pkg/router"
	"github.com/zhulinwei/go-dc/pkg/rpc"
	"log"
	"net/http"
	"time"
)

const (
	DebugMode      = "debug"
	ReleaseMode    = "release"
	ProductionMode = "production"
)

func main() {
	go rpc.GRPCRun(config.ServerConfig().GrpcPort)

	mode := config.ServerConfig().Mode
	if mode == ProductionMode || mode == ReleaseMode {
		gin.SetMode(gin.ReleaseMode)

	}
	route := gin.New()
	route.Use(gin.Logger())
	router.InitRoute(route)
	server := &http.Server{
		// 监听的TCP地址
		Addr: config.ServerConfig().HttpPort,
		// http句柄，用于处理程序响应的HTTP请求
		Handler: route,
		// 等待的最大时间
		IdleTimeout: 6 * time.Minute,
		// 允许读取的最大时间
		ReadTimeout: 30 * time.Second,
		// 允许写入的最大时间
		WriteTimeout: 30 * time.Second,
		// 请求头的最大字节数
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server run failed: %v", err)
	}
}
