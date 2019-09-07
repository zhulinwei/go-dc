package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/model"
	"github.com/zhulinwei/gin-demo/pkg/service"
)

type IUserController interface {
	Ping(ctx *gin.Context)
	SaveUser(ctx *gin.Context)
	QueryUserByName(ctx *gin.Context)
	UpdateUserByName(ctx *gin.Context)
	RemoveUserByName(ctx *gin.Context)
}

type UserController struct {
	userService service.IUserService
}

func BuildUserController () IUserController {
	return UserController{
		userService: service.BuildUserService(),
	}
}

func NewUserController(testService service.IUserService) IUserController {
	return UserController{
		userService: testService,
	}
}

// Ping
func (testController UserController) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

// Create
func (ctrl UserController) SaveUser(ctx *gin.Context) {
	// 解析前端数据
	var test model.UserReq
	if ctx.ShouldBind(&test) != nil {
		// TODO 错误处理
	}
	// 调用服务层逻辑
	saveID := ctrl.userService.SaveUser(test)
	// 返回处理结果
	ctx.JSON(http.StatusOK, gin.H{
		"id": saveID,
	})
}

// Read
func (ctrl UserController) QueryUserByName(ctx *gin.Context) {
	name := ctx.Param("name")
	result := ctrl.userService.QueryUserByName(name)
	ctx.JSON(http.StatusOK, result)
}

// Update
func (ctrl UserController) UpdateUserByName(ctx *gin.Context) {
	oldName := ctx.Param("name")

	var test model.UserReq
	if ctx.ShouldBind(&test) != nil {
		// TODO 错误处理
	}
	updateCount := ctrl.userService.UpdateUserByName(oldName, test.Name)
	ctx.JSON(http.StatusOK, gin.H{
		"updateCount": updateCount,
	})
}

// Delete
func (ctrl UserController) RemoveUserByName(ctx *gin.Context) {
	name := ctx.Param("name")

	DeletedCount := ctrl.userService.RemoveUserByName(name)
	ctx.JSON(http.StatusOK, gin.H{
		"DeletedCount": DeletedCount,
	})
}
