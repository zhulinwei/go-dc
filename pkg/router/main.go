package router

import "github.com/gin-gonic/gin"


func InitRouter () *gin.Engine {
	router := gin.Default()

	InitTest1Router(router)

	test2Router := new(Test2Router)
	test2Router.InitTest2Router(router)

	return router
}