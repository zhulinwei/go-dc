package router

import (
	"github.com/gin-gonic/gin"
)

func BuildRoute () *gin.Engine {
	route := gin.Default()

	BuildUserRouter().InitRouter(route)

	return route
}