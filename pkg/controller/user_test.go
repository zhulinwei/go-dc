package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zhulinwei/go-dc/pkg/model"
	mockService "github.com/zhulinwei/go-dc/pkg/service/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserController_QueryUserByName(t *testing.T) {
	//mock data
	const mockUrl = "/:name"
	const mockName = "tony"
	const mockMethod = "GET"
	mockObjectId := primitive.NewObjectID()

	// mock request
	route := gin.Default()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserService := mockService.NewMockIUserService(mockCtrl)
	mockUserService.EXPECT().QueryUserByName(mockName).Return(&model.UserDB{ID: mockObjectId, Age: 18, Name: mockName})

	mockUserController := UserController{
		userService: mockUserService,
	}
	route.GET(mockUrl, mockUserController.QueryUserByName)
	request := httptest.NewRequest(mockMethod, "/tony", nil)
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, request)

	body, err := ioutil.ReadAll(recorder.Result().Body)
	assert.NoError(t, err)

	var result model.UserDB
	err = json.Unmarshal(body, &result)
	assert.NoError(t, err)

	// assert result
	assert.NoError(t, err)
	assert.Equal(t, result.Name, mockName)
}
