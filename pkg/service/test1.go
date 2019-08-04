package service

import (
	"github.com/zhulinwei/gin-demo/pkg/dao"
	"github.com/zhulinwei/gin-demo/pkg/dto"
	"github.com/zhulinwei/gin-demo/pkg/model"
)

type Test1Service struct{}

func (*Test1Service) Ping() string {
	return "test1 service pong"
}

func (*Test1Service) SaveUser(test1 dto.User) interface{} {
	result := dao.GetTest1Dao().SaveUser(test1)
	return result.InsertedID
}

func (*Test1Service) QueryUserByName (name string) model.User {
	result := dao.GetTest1Dao().QueryUserByName(name)
	return result
}

func (*Test1Service) UpdateUserByName (oldName, newName string) interface{} {
	result := dao.GetTest1Dao().UpdateUserByName(oldName, newName)
	return result.ModifiedCount
}

func (*Test1Service) RemoveUserByName (name string) interface{} {
	result := dao.GetTest1Dao().RemoveUserByName(name)
	return result.DeletedCount
}