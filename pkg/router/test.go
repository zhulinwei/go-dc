package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/controller"
)

type TestRouter struct {
	TestController controller.ITestController
}

func NewTestRouter (testController controller.ITestController) *TestRouter {
	return &TestRouter{
		TestController: testController,
	}
}

func (testRouter *TestRouter) InitRouter(r *gin.Engine) {
	route := r.Group("/test1")

	route.GET("/ping", testRouter.TestController.Ping)
	route.POST("/users", testRouter.TestController.SaveUser)
	route.GET("/users/:name", testRouter.TestController.QueryUserByName)
	route.PUT("/users/:name", testRouter.TestController.UpdateUserByName)
	route.DELETE("/users/:name", testRouter.TestController.RemoveUserByName)
}
