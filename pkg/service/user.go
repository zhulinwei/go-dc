package service

import (
	"github.com/zhulinwei/gin-demo/pkg/dao"
	"github.com/zhulinwei/gin-demo/pkg/model"
)

type IUserService interface {
	SaveUser(test1 model.UserReq) interface{}
	QueryUserByName(name string) model.User
	RemoveUserByName(name string) interface{}
	UpdateUserByName(oldName, newName string) interface{}
}

type UserService struct {
	UserDao dao.IUserDao
}

func NewUserService (userDao dao.IUserDao) IUserService {
	return UserService{
		UserDao: userDao,
	}
}

func (service UserService) SaveUser (test model.UserReq) interface{} {
	result := service.UserDao.SaveUser(test)
	return result.InsertedID
}

func (service UserService) QueryUserByName(name string) model.User {
	result := service.UserDao.QueryUserByName(name)
	return result
}

func (service UserService) UpdateUserByName(oldName, newName string) interface{} {
	result := service.UserDao.UpdateUserByName(oldName, newName)
	return result.ModifiedCount
}

func (service UserService) RemoveUserByName(name string) interface{} {
	result := service.UserDao.RemoveUserByName(name)
	return result.DeletedCount
}
