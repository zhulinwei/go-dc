package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/controller"
)

func InitTest1Router(r *gin.Engine) {

	test1Controller := new(controller.Test1Controller)

	router := r.Group("/test1")
	router.GET("/ping", test1Controller.Ping)
}
