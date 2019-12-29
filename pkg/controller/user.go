package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/go-dc/pkg/model"
	"github.com/zhulinwei/go-dc/pkg/service"
	"github.com/zhulinwei/go-dc/pkg/util/log"
)

type IUserController interface {
	Ping(ctx *gin.Context)
	SaveUser(ctx *gin.Context)
	BulkSaveUser(ctx *gin.Context)
	QueryUserByName(ctx *gin.Context)
	UpdateUserByName(ctx *gin.Context)
	RemoveUserByName(ctx *gin.Context)
}

type UserController struct {
	userService service.IUserService
}

func BuildUserController() IUserController {
	return UserController{
		userService: service.BuildUserService(),
	}
}

// Ping
func (UserController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// Create
func (ctrl UserController) SaveUser(ctx *gin.Context) {
	// 解析前端数据
	var user model.UserRequest
	if err := ctx.ShouldBind(&user); err != nil {
		log.Error("gin bind user error", log.String("error", err.Error()))
		return
	}
	// 调用服务层逻辑
	saveID := ctrl.userService.SaveUser(user)
	// 返回处理结果
	ctx.JSON(http.StatusOK, gin.H{"id": saveID})
}

func (ctrl UserController) BulkSaveUser (ctx *gin.Context) {
	var users []model.UserRequest
	if err := ctx.ShouldBind(&users); err != nil {
		log.Error("gin bind users error", log.String("error", err.Error()))
		return
	}
	log.Info("gin bind users success", log.Reflect("users", users))
	// 调用服务层逻辑
	saveID := ctrl.userService.BulkSaveUser(users)
	// 返回处理结果
	ctx.JSON(http.StatusOK, gin.H{"id": saveID})
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

	var test model.UserRequest
	if err := ctx.ShouldBind(&test); err != nil {
		log.Error("gin bind error", log.String("error", err.Error()))
		return
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
