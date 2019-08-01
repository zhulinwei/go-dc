package router

import "github.com/gin-gonic/gin"


func InitRouter () *gin.Engine {
	router := gin.Default()

	InitTest1Router(router)
	InitTest2Router(router)

	return router
}