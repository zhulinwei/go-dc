package service

import (
	"github.com/zhulinwei/go-dc/pkg/cache"
	"github.com/zhulinwei/go-dc/pkg/dao"
	"github.com/zhulinwei/go-dc/pkg/model"
)

type IUserService interface {
	SaveUser(user model.UserRequest) error
	BulkSaveUser(users []model.UserRequest) error
	QueryUserByName(name string) (*model.UserDB, error)
	QueryUsersByName(name string) ([]model.UserDB, error)
	RemoveUserByName(name string) error
	UpdateUserByName(oldName, newName string) error
	SaveUserAmount(userAmount model.UserAmountRequest) error
}

type UserService struct {
	Cache     cache.ICache
	UserDao   dao.IUserDao
	AmountDao dao.IAmountDao
}

func BuildUserService() IUserService {
	return UserService{
		Cache:     cache.BuildCache(),
		UserDao:   dao.BuildUserDao(),
		AmountDao: dao.BuildAmountDao(),
	}
}

// save user
func (service UserService) SaveUser(user model.UserRequest) error {
	_, err := service.UserDao.SaveUser(user)
	return err
}

// bulk save user
func (service UserService) BulkSaveUser(users []model.UserRequest) error {
	_, err := service.UserDao.BulkSaveUser(users)
	return err
}

func (service UserService) QueryUserByName(name string) (*model.UserDB, error) {
	return service.UserDao.QueryUserByName(name)
}

func (service UserService) QueryUsersByName(name string) ([]model.UserDB, error) {
	return service.UserDao.QueryUsersByName(name)
}

func (service UserService) UpdateUserByName(oldName, newName string) error {
	_, err := service.UserDao.UpdateUserByName(oldName, newName)
	return err
}

func (service UserService) RemoveUserByName(name string) error {
	_, err := service.UserDao.RemoveUserByName(name)
	return err
}

func (service UserService) SaveUserAmount(userAmount model.UserAmountRequest) error {
	_, err := service.AmountDao.SaveAmount(userAmount)
	return err
}
