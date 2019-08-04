package controller

import (
"github.com/gin-gonic/gin"
"github.com/zhulinwei/gin-demo/pkg/service"
)

type Test2Controller struct{}

func (*Test2Controller) Ping(context *gin.Context) {
	test2Service := new(service.Test2Service)
	context.JSON(200, gin.H{
		"message": test2Service.Ping(),
	})
}
