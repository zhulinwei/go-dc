package main

import "github.com/zhulinwei/gin-demo/pkg/router"


func main () {
	r := router.InitRouter()
	r.Run()
}
