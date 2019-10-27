package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zhulinwei/go-dc/pkg/model"
	mockService "github.com/zhulinwei/go-dc/pkg/service/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestUserController_QueryUserByName(t *testing.T) {
	//mock data
	const mockName = "tony"
	const mockMethod = "GET"
	const mockUrl = "/test1/users/tony"
	mockObjectId := primitive.NewObjectID()

	// mock request
	route := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserService := mockService.NewMockIUserService(mockCtrl)
	mockUserService.EXPECT().QueryUserByName(mockName).Return(model.User{ Test1ID: mockObjectId, Age: 18, Name: mockName})

	mockUserController := UserController{
		userService:mockUserService,
	}
	route.GET(mockUrl, mockUserController.QueryUserByName)
	request := httptest.NewRequest(mockMethod, mockUrl, nil)
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)

	body, err := ioutil.ReadAll(recorder.Result().Body)
	realResult, err := strconv.Unquote(string(body))

	// assert result
	assert.NoError(t, err)
	assert.Equal(t, realResult, mockName)
}
