package controller

import "github.com/zhulinwei/gin-demo/pkg/service"

func InitController(userService service.IUserService) *UserController {
	return NewUserController(userService)
}
