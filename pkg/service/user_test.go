package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mockDao "github.com/zhulinwei/go-dc/pkg/dao/mock"
	"github.com/zhulinwei/go-dc/pkg/model"
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
	mockUserService := UserService{
		UserDao: mockUserDao,
	}
	realResult := mockUserService.SaveUser(mockTest)

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
	mockUserDao := mockDao.NewMockIUserDao(mockCtrl)
	mockUserService := UserService{
		UserDao: mockUserDao,
	}
	mockUserDao.EXPECT().QueryUserByName(mockName).Return(model.UserDB{Test1ID: mockObjectId, Age: 18, Name: "tony"})
	realResult := mockUserService.QueryUserByName(mockName)

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
	mockUserDao := mockDao.NewMockIUserDao(mockCtrl)
	mockUserService := UserService{
		UserDao: mockUserDao,
	}
	mockUserDao.EXPECT().UpdateUserByName(mockOldName, mockNewName).Return(&mongo.UpdateResult{ModifiedCount: mockCount})
	realResult := mockUserService.UpdateUserByName(mockOldName, mockNewName)

	// assert result
	assert.Equal(t, mockCount, realResult)
}

func TestTestService_RemoveUserByName(t *testing.T) {
	const mockName = "tony"
	const mockCount = int64(1)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserDao := mockDao.NewMockIUserDao(mockCtrl)
	mockUserService := UserService{
		UserDao: mockUserDao,
	}
	mockUserDao.EXPECT().RemoveUserByName(mockName).Return(&mongo.DeleteResult{DeletedCount: mockCount})
	realResult := mockUserService.RemoveUserByName(mockName)

	// assert result
	assert.Equal(t, mockCount, realResult)
}
