package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter (route *gin.Engine){
	userRouter := BuildUserRouter()

	userRouter.InitRouter(route)
}

