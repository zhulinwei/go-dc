package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/controller"
	"github.com/zhulinwei/gin-demo/pkg/dao"
	"github.com/zhulinwei/gin-demo/pkg/database"
	"github.com/zhulinwei/gin-demo/pkg/router"
	"github.com/zhulinwei/gin-demo/pkg/service"
)


func main () {
	route := gin.Default()
	// 初始化数据库
	cache, mongo := database.InitDatabase()

	// 初始化数据交互层
	testDao := dao.InitDao(mongo)

	// 初始化服务层
	testService := service.InitService(testDao)

	// 初始化控制层
	testController := controller.InitController(testService)

	// 初始化路由层
	router.InitRouter(route, testController)

	fmt.Println(cache, mongo, testDao, testService, testController)
	route.Run()
}