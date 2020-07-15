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
	if result, err := service.UserDao.SaveUser(user); err != nil {
		return nil
	} else {
		return result.InsertedID
	}
}

// bulk save user
func (service UserService) BulkSaveUser(users []model.UserRequest) int64 {
	result, err := service.UserDao.BulkSaveUser(users)
	if err != nil {
		return 0
	}
	return result.InsertedCount
}

// query user
func (service UserService) QueryUserByName(name string) (*model.UserDB, error) {
	return service.UserDao.QueryUserByName(name)
}

func (service UserService) QueryUsersByName(name string) ([]model.UserDB, error) {
	return service.UserDao.QueryUsersByName(name)
}

func (service UserService) UpdateUserByName(oldName, newName string) interface{} {
	result, err := service.UserDao.UpdateUserByName(oldName, newName)
	if err != nil {
		return nil
	}

	return result.ModifiedCount
}

func (service UserService) RemoveUserByName(name string) interface{} {
	result, err := service.UserDao.RemoveUserByName(name)
	if err != nil {
		return nil
	}

	return result.DeletedCount
}
