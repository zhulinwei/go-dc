package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/dto"
	"github.com/zhulinwei/gin-demo/pkg/service"
	"net/http"
)

type ITestController interface {
	Ping(ctx *gin.Context)
	SaveUser(ctx *gin.Context)
	QueryUserByName(ctx *gin.Context)
	UpdateUserByName(ctx *gin.Context)
	RemoveUserByName(ctx *gin.Context)
}

type TestController struct {
	TestService service.ITestService
}

func NewTestController (testService service.ITestService) *TestController {
	return &TestController{
		TestService: testService,
	}
}

// Ping
func (testController *TestController) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

// Create
func (testController *TestController) SaveUser(ctx *gin.Context) {
	// 解析前端数据
	var test dto.Test
	if ctx.ShouldBind(&test) != nil {
		// TODO 错误处理
	}
	// 调用服务层逻辑
	saveID := testController.TestService.SaveUser(test)
	// 返回处理结果
	ctx.JSON(http.StatusOK, gin.H{
		"id": saveID,
	})
}

// Read
func (testController *TestController) QueryUserByName(ctx *gin.Context) {
	name := ctx.Param("name")
	result := testController.TestService.QueryUserByName(name)
	ctx.JSON(http.StatusOK, result)
}

// Update
func (testController *TestController) UpdateUserByName(ctx *gin.Context) {
	oldName := ctx.Param("name")

	var test dto.Test
	if ctx.ShouldBind(&test) != nil {
		// TODO 错误处理
	}
	updateCount := testController.TestService.UpdateUserByName(oldName, test.Name)
	ctx.JSON(http.StatusOK, gin.H{
		"updateCount": updateCount,
	})
}

// Delete
func (testController *TestController) RemoveUserByName(ctx *gin.Context) {
	name := ctx.Param("name")

	DeletedCount := testController.TestService.RemoveUserByName(name)
	ctx.JSON(http.StatusOK, gin.H{
		"DeletedCount": DeletedCount,
	})
}