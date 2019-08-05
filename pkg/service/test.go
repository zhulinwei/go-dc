package service

import (
	"github.com/zhulinwei/gin-demo/pkg/dao"
	"github.com/zhulinwei/gin-demo/pkg/dto"
	"github.com/zhulinwei/gin-demo/pkg/model"
)

type ITestService interface {
	SaveUser(test1 dto.Test) interface{}
	QueryUserByName(name string) model.Test
	RemoveUserByName(name string) interface{}
	UpdateUserByName(oldName, newName string) interface{}
}

type TestService struct {
	TestDao dao.ITestDao
}

func NewTestService (testDao dao.ITestDao) *TestService {
	return &TestService{
		TestDao: testDao,
	}
}

func (service *TestService) SaveUser (test dto.Test) interface{} {
	result := service.TestDao.SaveUser(test)
	return result.InsertedID
}

func (service *TestService) QueryUserByName(name string) model.Test {
	result := service.TestDao.QueryUserByName(name)
	return result
}

func (service *TestService) UpdateUserByName(oldName, newName string) interface{} {
	result := service.TestDao.UpdateUserByName(oldName, newName)
	return result.ModifiedCount
}

func (service *TestService) RemoveUserByName(name string) interface{} {
	result := service.TestDao.RemoveUserByName(name)
	return result.DeletedCount
}
