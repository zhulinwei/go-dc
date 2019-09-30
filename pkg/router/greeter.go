package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/go-dc/pkg/controller"
)

type IGreeterRouter interface {
	InitRouter(r *gin.Engine)
}

type GreeterRouter struct {
	GreeterController controller.IGreeterController
}

func BuildGreeterRouter() IGreeterRouter {
	return GreeterRouter{
		GreeterController: controller.BuildGreeterController(),
	}
}

func (greeterRouter GreeterRouter) InitRouter(r *gin.Engine) {
	route := r.Group("/grpc")

	route.GET("/greeter/:name", greeterRouter.GreeterController.QueryGreeterFromGrpc)
}
