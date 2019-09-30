package service

import (
	"github.com/zhulinwei/go-dc/pkg/cache"
	"github.com/zhulinwei/go-dc/pkg/dao"
	"github.com/zhulinwei/go-dc/pkg/model"
)

type IUserService interface {
	SaveUser(test1 model.UserReq) interface{}
	QueryUserByName(name string) model.User
	RemoveUserByName(name string) interface{}
	UpdateUserByName(oldName, newName string) interface{}
}

type UserService struct {
	Cache   cache.ICache
	UserDao dao.IUserDao
}

func BuildUserService() IUserService {
	return UserService{
		Cache:   cache.BuildCache(),
		UserDao: dao.BuildUserDao(),
	}
}

func (service UserService) SaveUser(test model.UserReq) interface{} {
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
