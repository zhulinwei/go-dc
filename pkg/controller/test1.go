package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/service"
)

type Test1Controller struct{}

func (ctrl *Test1Controller) Ping(context *gin.Context) {
	test := new(service.Test1Service)
	context.JSON(200, gin.H{
		"message": test.Ping(),
	})
}
