package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/controller"
)

type Test1Router struct{}

func (*Test1Router) InitRouter(r *gin.Engine) {
	test1Controller := new(controller.Test1Controller)

	router := r.Group("/test1")
	router.GET("/ping", test1Controller.Ping)
}

//var test1Controller = new(controller.Test1Controller)
//
//func InitTest1Router(r *gin.Engine) {
//
//	router := r.Group("/test1")
//	router.GET("/ping", test1Controller.Ping)
//}
//
//
//type Test1Controller struct{}
//
//func (ctrl *Test1Controller) Ping(context *gin.Context) {
//	test1Service := new(service.Test1Service)
//	context.JSON(200, gin.H{
//		"message": test1Service.Ping(),
//	})
//}
