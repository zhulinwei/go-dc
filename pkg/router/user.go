package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/controller"
)

type UserRouter struct {
	UserController controller.IUserController
}

func NewUserRouter (userController controller.IUserController) *UserRouter {
	return &UserRouter{
		UserController: userController,
	}
}

func (userRouter *UserRouter) InitRouter(r *gin.Engine) {
	route := r.Group("/test1")

	route.GET("/ping", userRouter.UserController.Ping)
	route.POST("/users", userRouter.UserController.SaveUser)
	route.GET("/users/:name", userRouter.UserController.QueryUserByName)
	route.PUT("/users/:name", userRouter.UserController.UpdateUserByName)
	route.DELETE("/users/:name", userRouter.UserController.RemoveUserByName)
}
