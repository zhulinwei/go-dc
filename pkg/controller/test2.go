package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/service"
)

type Test2Controller struct{
	Test2Service service.Test2Service
}

func (ctrl *Test2Controller) Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": ctrl.Test2Service.Ping(),
	})
}
