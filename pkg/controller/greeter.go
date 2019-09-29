package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/service"
	greeterPb "github.com/zhulinwei/grpc-demo/helloworld/greeter/proto"
	"net/http"
)

type IGreeterController interface {
	QueryGreeterFromGrpc(ctx *gin.Context)
}

type GreeterController struct {
	greeterService service.IGreeterClient
}

func BuildGreeterController() IGreeterController {
	return GreeterController{
		greeterService: service.BuildGreeterService(),
	}
}

// Query greeter from other grpc
func (ctrl GreeterController) QueryGreeterFromGrpc(ctx *gin.Context) {
	name := ctx.Param("name")

	var err error
	var response *greeterPb.HelloReply
	if response, err = ctrl.greeterService.QueryGreeterFromGrpc(name); err != nil {
		// TODO 错误处理
	}

	var message string
	if response != nil {
		message = response.Message
	} else {
		message = "nothing"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
