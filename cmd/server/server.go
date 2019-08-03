package main

import (
	_ "github.com/zhulinwei/gin-demo/pkg/dao"
	"github.com/zhulinwei/gin-demo/pkg/router"
)

func main() {
	route := router.GetRoute()
	route.Run()
}
