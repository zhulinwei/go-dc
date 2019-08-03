package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/controller"
)

type Test2Router struct{}

func (*Test2Router) InitRouter(r *gin.Engine) {
	test2Controller := new(controller.Test2Controller)

	router := r.Group("/test2")
	router.GET("/ping", test2Controller.Ping)
}