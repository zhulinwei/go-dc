package main

import (
	"github.com/zhulinwei/gin-demo/pkg/database"
	"github.com/zhulinwei/gin-demo/pkg/router"
)

func main() {
	db := database.InitDatabase()

	r := router.InitRouter()
	r.Run()
}
