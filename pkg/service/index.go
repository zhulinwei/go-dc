package service

import "github.com/zhulinwei/gin-demo/pkg/dao"

func InitService (testDao dao.IUserDao) *UserService {
	return NewUserService(testDao)
}