package controller

import (
	"github.com/gin-gonic/gin"
)

type Test1Controller struct{}

func (ctrl *Test1Controller) Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "test1 pong",
	})
}
