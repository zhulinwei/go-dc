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

	router.POST("/users", test1Controller.SaveUser)
	router.GET("/users/:name", test1Controller.QueryUserByName)
	router.PUT("/users/:name", test1Controller.UpdateUserByName)
	router.DELETE("/users/:name", test1Controller.RemoveUserByName)

}
