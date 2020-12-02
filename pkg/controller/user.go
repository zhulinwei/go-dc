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
	SaveUserAmount(ctx *gin.Context)
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
	if err := ctrl.userService.SaveUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{Code: 1001, Msg: "save user fail"})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 0, Msg: "success"})
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
	if err := ctrl.userService.BulkSaveUser(usersRequest.Users); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{Code: 1001, Msg: "bulk save user fail"})
		return
	}
	// 返回处理结果
	ctx.JSON(http.StatusOK, model.Response{Code: 0, Msg: "success"})
}

// Read And Return Single User
func (ctrl UserController) QueryUserByName(ctx *gin.Context) {
	name := ctx.Param("name")
	user, err := ctrl.userService.QueryUserByName(name)
	if err != nil {
		log.Error("query user fail", log.String("error", err.Error()))
		ctx.JSON(http.StatusInternalServerError, model.Response{Code: 1001, Msg: "query user fail"})
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
		ctx.JSON(http.StatusInternalServerError, model.Response{Code: 1001, Msg: "query user fail"})
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
	if err := ctrl.userService.UpdateUserByName(oldName, user.Name); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{Code: 1001, Msg: "update user fail"})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Code: 0, Msg: "success"})
}

// Delete
func (ctrl UserController) RemoveUserByName(ctx *gin.Context) {
	name := ctx.Param("name")

	if err := ctrl.userService.RemoveUserByName(name); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{Code: 1001, Msg: "remove user fail"})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{Code: 0, Msg: "success"})
}

// Create User Amount
func (ctrl UserController) SaveUserAmount(ctx *gin.Context) {
	// 解析前端数据
	var userAmount model.UserAmountRequest
	if err := ctx.ShouldBind(&userAmount); err != nil {
		log.Error("gin bind user amount error", log.String("error", err.Error()))
		ctx.JSON(http.StatusBadRequest, model.Response{Code: -1, Msg: util.ParserErrorMsg(err)})
		return
	}
	// 调用服务层逻辑
	if err := ctrl.userService.SaveUserAmount(userAmount); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{Code: 1001, Msg: "save user amount fail"})
		return
	}
	// 返回处理结果
	ctx.JSON(http.StatusOK, model.Response{Code: 0, Msg: "success"})
}
