package router

import "github.com/gin-gonic/gin"

func Default () *gin.Engine {
	engine := gin.New()

	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{ "message": "pong" })
	})

	return engine
}