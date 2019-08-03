package router

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()
var test1Router = new(Test1Router)
var test2Router = new(Test2Router)

func GetRoute () *gin.Engine {
	return router
}

func init() {
	test1Router.InitRouter(router)
	test2Router.InitRouter(router)
}
