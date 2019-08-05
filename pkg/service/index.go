package service

import "github.com/zhulinwei/gin-demo/pkg/dao"

func InitService (testDao dao.ITestDao) *TestService {
	return NewTestService(testDao)
}