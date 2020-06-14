package service

import (
	"github.com/zhulinwei/go-dc/pkg/cache"
	"github.com/zhulinwei/go-dc/pkg/dao"
	"github.com/zhulinwei/go-dc/pkg/model"
)

type IUserService interface {
	SaveUser(user model.UserRequest) interface{}
	BulkSaveUser(users []model.UserRequest) int64
	QueryUserByName(name string) (*model.UserDB, error)
	QueryUsersByName(name string) ([]model.UserDB, error)
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

// save user
func (service UserService) SaveUser(user model.UserRequest) interface{} {
	return service.UserDao.SaveUser(user).InsertedID
}

// bulk save user
func (service UserService) BulkSaveUser(users []model.UserRequest) int64 {
	return service.UserDao.BulkSaveUser(users).InsertedCount
}

// query user
func (service UserService) QueryUserByName(name string) (*model.UserDB, error) {
	return service.UserDao.QueryUserByName(name)
}

func (service UserService) QueryUsersByName(name string) ([]model.UserDB, error) {
	return service.UserDao.QueryUsersByName(name)
}

func (service UserService) UpdateUserByName(oldName, newName string) interface{} {
	return service.UserDao.UpdateUserByName(oldName, newName).ModifiedCount
}

func (service UserService) RemoveUserByName(name string) interface{} {
	return service.UserDao.RemoveUserByName(name).DeletedCount
}
