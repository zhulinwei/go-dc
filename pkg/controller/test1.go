package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhulinwei/gin-demo/pkg/service"
)

type Test1Controller struct{}

var test1Service = new(service.Test1Service)

func (*Test1Controller) Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": test1Service.Ping(),
	})
}

type User struct {
	Age  int    `json:"age" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func (*Test1Controller) SaveUser(context *gin.Context) {
	var user User
	if context.ShouldBind(&user) != nil {
		// TODO 错误处理
		fmt.Println("something wrong")
	}
	fmt.Println(user)

	//newUser := &model.Test1Model{
	//	Age:  18,
	//	Name: "tony",
	//}
	//context.JSON(200, gin.H{
	//	"message": test1Service.Ping(),
	//})
	//test1Service.SaveUser(newUser)
}
