//package main
//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/zhulinwei/gin-demo/pkg/controller"
//	"github.com/zhulinwei/gin-demo/pkg/dao"
//	"github.com/zhulinwei/gin-demo/pkg/database"
//	"github.com/zhulinwei/gin-demo/pkg/router"
//	"github.com/zhulinwei/gin-demo/pkg/service"
//)
//
//
//func main () {
//	route := gin.Default()
//	// 初始化数据库
//	cache, mongo := database.InitDatabase()
//
//	// 初始化数据交互层
//	userDao := dao.InitDao(mongo)
//
//	// 初始化服务层
//	userService := service.InitService(userDao)
//
//	// 初始化控制层
//	userController := controller.InitController(userService)
//
//	// 初始化路由层
//	router.InitRouter(route, userController)
//
//	fmt.Println(cache, mongo, userDao, userController, userService)
//	route.Run()
//}

package main

/*
#include <stdio.h>

void printint(int v) {
    printf("printint: %d\n", v);
}
*/
import "C"

func main() {
	v := 42
	C.printint(C.int(v))
}