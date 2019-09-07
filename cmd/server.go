package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/router"
)

func main () {
	route := gin.Default()
	// 初始化数据库
	//cache, mongo := database.InitDatabase()

	// 初始化数据交互层
	//userDao := dao.InitDao(mongo)

	// 初始化服务层
	//userService := service.InitService(userDao)

	// 初始化控制层
	//userController := controller.InitController(userService)

	// 初始化路由层
	router.InitRouter(route)

	//fmt.Println(cache, mongo, userDao)
	err := route.Run()
	fmt.Println(err)
}

