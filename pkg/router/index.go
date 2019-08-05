package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/controller"
)

func InitRouter (route *gin.Engine, userController controller.IUserController){
	testRouter := NewTestRouter(userController)

	testRouter.InitRouter(route)
}

