package controller

import "github.com/zhulinwei/gin-demo/pkg/service"

func InitController (testService service.ITestService) *TestController {
	return NewTestController(testService)
}
