package router

import (
	"github.com/gin-gonic/gin"
)

func InitTest2Router (r *gin.Engine) {
	router := r.Group("/test2")
	router.GET("v1", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
}