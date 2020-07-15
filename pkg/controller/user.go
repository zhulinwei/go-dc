package controller

import (
	"net/http"

	"github.com/zhulinwei/go-dc/pkg/util"

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
	QueryUsersByName(ctx *gin.Context)
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
		ctx.JSON(http.StatusBadRequest, model.Response{Code: -1, Msg: util.ParserErrorMsg(err)})
		return
	}
	// 调用服务层逻辑
	saveID := ctrl.userService.SaveUser(user)
	// 返回处理结果
	ctx.JSON(http.StatusOK, model.Response{Code: 0, Msg: "success", Data: gin.H{"id": saveID}})
}

// Bulk Create
func (ctrl UserController) BulkSaveUser(ctx *gin.Context) {
	//	var users []model.UserRequest
	var usersRequest model.UsersRequest

	if err := ctx.ShouldBind(&usersRequest); err != nil {
		log.Error("gin bind users error", log.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, model.Response{Code: -1, Msg: util.ParserErrorMsg(err)})
		return
	}
	log.Info("gin bind users success", log.Reflect("users", usersRequest))
	// 调用服务层逻辑
	saveCount := ctrl.userService.BulkSaveUser(usersRequest.Users)

	// // 返回处理结果
	ctx.JSON(http.StatusOK, model.Response{Code: 0, Msg: "success", Data: gin.H{"saveCount": saveCount}})
}

// Read And Return Single User
func (ctrl UserController) QueryUserByName(ctx *gin.Context) {
	name := ctx.Param("name")
	// 可以进一步判断user是否为nil值
	user, err := ctrl.userService.QueryUserByName(name)
	if err != nil {
		log.Error("query user fail", log.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 0, Msg: "success", Data: user})
}

// Read And Return Multiple User
func (ctrl UserController) QueryUsersByName(ctx *gin.Context) {
	name := ctx.Param("name")
	users, err := ctrl.userService.QueryUsersByName(name)
	if err != nil {
		log.Error("query user fail", log.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 0, Msg: "success", Data: users})
}

// Update
func (ctrl UserController) UpdateUserByName(ctx *gin.Context) {
	oldName := ctx.Param("name")

	var user model.UserRequest
	if err := ctx.ShouldBind(&user); err != nil {
		log.Error("gin bind error", log.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, model.Response{Code: -1, Msg: util.ParserErrorMsg(err)})
		return
	}
	updateCount := ctrl.userService.UpdateUserByName(oldName, user.Name)

	ctx.JSON(http.StatusOK, model.Response{Code: 0, Msg: "success", Data: gin.H{"updateCount": updateCount}})
}

// Delete
func (ctrl UserController) RemoveUserByName(ctx *gin.Context) {
	name := ctx.Param("name")

	deletedCount := ctrl.userService.RemoveUserByName(name)
	ctx.JSON(http.StatusOK, model.Response{Code: 0, Msg: "success", Data: gin.H{"deletedCount": deletedCount}})
}
