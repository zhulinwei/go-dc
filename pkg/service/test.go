package service

import (
	"github.com/zhulinwei/gin-demo/pkg/dao"
	"github.com/zhulinwei/gin-demo/pkg/dto"
)

type ITestService interface {
	SaveUser(test1 dto.Test) interface{}
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