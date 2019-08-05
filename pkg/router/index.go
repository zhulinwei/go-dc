package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/controller"
)

func InitRouter (route *gin.Engine, testController controller.ITestController){
	testRouter := NewTestRouter(testController)

	testRouter.InitRouter(route)
}

