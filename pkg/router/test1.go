package router

import (
	"github.com/gin-gonic/gin"

	"github.com/zhulinwei/gin-demo/pkg/controller"
)

func InitTest1Router (r *gin.Engine) {
	router := r.Group("/test1")

	router.GET("/ping", controller.Test1Controller.ping)

	router.GET("v1", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
}