package service

import "github.com/zhulinwei/gin-demo/pkg/dao"

func InitService (testDao dao.IUserDao) IUserService {
	return NewUserService(testDao)
}