package router

import (
	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.Engine) {
	BuildUserRouter().InitRouter(route)
	BuildGreeterRouter().InitRouter(route)
}