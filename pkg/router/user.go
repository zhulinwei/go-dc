package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/go-dc/pkg/controller"
)

type IUserRouter interface {
	InitRouter(r *gin.Engine)
}

type UserRouter struct {
	UserController controller.IUserController
}

func BuildUserRouter() IUserRouter {
	return UserRouter{
		UserController: controller.BuildUserController(),
	}
}

func (userRouter UserRouter) InitRouter(r *gin.Engine) {
	route := r.Group("/api")

	route.GET("/ping", userRouter.UserController.Ping)
	// 单次保存
	route.POST("/v1/users", userRouter.UserController.SaveUser)
	// 批量保存
	route.POST("/v2/users", userRouter.UserController.BulkSaveUser)

	// 单个用户
	route.GET("/v1/users/:name", userRouter.UserController.QueryUserByName)
	route.GET("/v2/users/:name", userRouter.UserController.QueryUsersByName)
	route.PUT("/v1/users/:name", userRouter.UserController.UpdateUserByName)
	route.DELETE("/v1/users/:name", userRouter.UserController.RemoveUserByName)
}
