package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/dto"
	"github.com/zhulinwei/gin-demo/pkg/service"
	"net/http"
)

type Test1Controller struct{}

var test1Service = new(service.Test1Service)

func (*Test1Controller) Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": test1Service.Ping(),
	})
}

// Create
func (*Test1Controller) SaveUser(ctx *gin.Context) {
	// 解析前端数据
	var user dto.User
	if ctx.ShouldBind(&user) != nil {
		// TODO 错误处理
	}

	// 调用服务层逻辑
	saveID := test1Service.SaveUser(user)

	// 返回处理结果
	ctx.JSON(http.StatusOK, gin.H{
		"id": saveID,
	})
}

// Read
func (*Test1Controller) QueryUserByName(ctx *gin.Context) {
	name := ctx.Param("name")
	result := test1Service.QueryUserByName(name)
	ctx.JSON(http.StatusOK, result)
}

// Update
func (*Test1Controller) UpdateUserByName(ctx *gin.Context) {
	oldName := ctx.Param("name")

	var user dto.User
	if ctx.ShouldBind(&user) != nil {
		// TODO 错误处理
	}
	updateCount := test1Service.UpdateUserByName(oldName, user.Name)
	ctx.JSON(http.StatusOK, gin.H{
		"updateCount": updateCount,
	})
}

// Delete
func (*Test1Controller) RemoveUserByName(ctx *gin.Context) {
	name := ctx.Param("name")

	DeletedCount := test1Service.RemoveUserByName(name)
	ctx.JSON(http.StatusOK, gin.H{
		"DeletedCount": DeletedCount,
	})
}