package service

import (
	"github.com/zhulinwei/go-dc/pkg/cache"
	"github.com/zhulinwei/go-dc/pkg/dao"
	"github.com/zhulinwei/go-dc/pkg/model"
	"github.com/zhulinwei/go-dc/pkg/util/log"
)

type IUserService interface {
	SaveUser(user model.UserRequest) interface{}
	BulkSaveUser(users []model.UserRequest) interface{}
	QueryUserByName(name string) model.UserDB
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

func (service UserService) SaveUser(user model.UserRequest) interface{} {
	result := service.UserDao.SaveUser(user)
	return result.InsertedID
}

func (service UserService) BulkSaveUser(users []model.UserRequest) interface{} {

	docs := []interface{}{users}
	log.Debug("docs", log.Reflect("docs", docs))

	result := service.UserDao.BulkSaveUser(users)

	return result.InsertedIDs
}

func (service UserService) QueryUserByName(name string) model.UserDB {
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
