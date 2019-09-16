package main

import (
	"fmt"
	"github.com/zhulinwei/gin-demo/pkg/router"
)

func main () {
	route := router.BuildRoute()

	if err := route.Run(); err != nil {
		fmt.Print(err)
	}
}
