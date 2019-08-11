package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mockDao "github.com/zhulinwei/gin-demo/pkg/dao/mock"
	"github.com/zhulinwei/gin-demo/pkg/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestTestService_SaveUser(t *testing.T) {
	// mock data
	mockTest := model.UserReq{Age: 18, Name: "tony"}
	mockObjectId := primitive.NewObjectID()

	// mock request
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserDao := mockDao.NewMockIUserDao(mockCtrl)
	mockUserDao.EXPECT().SaveUser(mockTest).Return(&mongo.InsertOneResult{InsertedID: mockObjectId})
	mockTestService := NewUserService(mockUserDao)
	realResult := mockTestService.SaveUser(mockTest)

	// assert result
	assert.Equal(t, mockObjectId, realResult)
}

func TestTestService_QueryUserByName(t *testing.T) {
	// mock data
	const mockName = "tony"
	mockObjectId := primitive.NewObjectID()

	// mock request
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockTestDao := mockDao.NewMockIUserDao(mockCtrl)
	mockTestService := NewUserService(mockTestDao)
	mockTestDao.EXPECT().QueryUserByName(mockName).Return(model.User{Test1ID: mockObjectId, Age: 18, Name: "tony"})
	realResult := mockTestService.QueryUserByName(mockName)

	// assert result
	assert.Equal(t, mockName, realResult.Name)
}

func TestTestService_UpdateUserByName(t *testing.T) {
	// mock data
	const mockCount = int64(1)
	const mockOldName = "tony1"
	const mockNewName = "tony2"

	// mock request
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockTestDao := mockDao.NewMockIUserDao(mockCtrl)
	mockTestService := NewUserService(mockTestDao)
	mockTestDao.EXPECT().UpdateUserByName(mockOldName, mockNewName).Return(&mongo.UpdateResult{ModifiedCount: mockCount})
	realResult := mockTestService.UpdateUserByName(mockOldName, mockNewName)

	// assert result
	assert.Equal(t, mockCount, realResult)
}

func TestTestService_RemoveUserByName(t *testing.T) {
	const mockName = "tony"
	const mockCount = int64(1)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockTestDao := mockDao.NewMockIUserDao(mockCtrl)
	mockTestService := NewUserService(mockTestDao)
	mockTestDao.EXPECT().RemoveUserByName(mockName).Return(&mongo.DeleteResult{DeletedCount: mockCount})
	realResult := mockTestService.RemoveUserByName(mockName)

	// assert result
	assert.Equal(t, mockCount, realResult)
}
