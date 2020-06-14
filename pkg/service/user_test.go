package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mockDao "github.com/zhulinwei/go-dc/pkg/dao/mock"
	"github.com/zhulinwei/go-dc/pkg/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestUserService_SaveUser(t *testing.T) {
	// mock data
	mockTest := model.UserRequest{Age: 18, Name: "tony"}
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

func TestUserService_QueryUserByName(t *testing.T) {
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
	mockUserDao.EXPECT().QueryUserByName(mockName).Return(&model.UserDB{ID: mockObjectId, Age: 18, Name: "tony"}, nil)
	realResult, err := mockUserService.QueryUserByName(mockName)

	// assert result
	assert.NoError(t, err)
	assert.Equal(t, mockName, realResult.Name)
}

func TestUserService_UpdateUserByName(t *testing.T) {
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

func TestUserService_RemoveUserByName(t *testing.T) {
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
