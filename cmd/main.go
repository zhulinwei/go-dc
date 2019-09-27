package main

import (
	_ "github.com/zhulinwei/gin-demo/pkg/config"
	"github.com/zhulinwei/gin-demo/pkg/router"
	"log"
)

func main() {
	if err := router.BuildRoute().Run(); err != nil {
		log.Fatalf("server run failed: %v", err)
	}
}
