package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/controller"
)

type Test2Router struct {
	Test2Controller controller.Test2Controller
}



func (router Test2Router) InitTest2Router(r *gin.Engine) {

	//test2Controller := new(controller.Test2Controller)
	route := r.Group("/test2")
	//router.GET("/ping", Test2Router.Test2Controller.Ping)
	route.GET("/ping", router.Test2Controller.Ping)
}
