package dao

import (
	"github.com/zhulinwei/gin-demo/pkg/model"
	"github.com/zhulinwei/gin-demo/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserDao interface {
	QueryUserByName(name string) model.User
	SaveUser(test1 model.UserReq) *mongo.InsertOneResult
	RemoveUserByName(name string) *mongo.DeleteResult
	UpdateUserByName(oldName, newName string) *mongo.UpdateResult
}

type UserDao struct {
	UserCollection *mongo.Collection
}

func NewUserDao (mongodb *mongo.Client) IUserDao {
	return UserDao{
		UserCollection: mongodb.Database("test_database").Collection("test_collection"),
	}
}

func (userDao UserDao) SaveUser(test1 model.UserReq) *mongo.InsertOneResult {
	var err error
	var result *mongo.InsertOneResult
	if result, err = userDao.UserCollection.InsertOne(util.GetHelper().GetContent(), test1); err != nil {
		// TODO 错误处理
	}
	return result
}

func (userDao UserDao) QueryUserByName(name string) model.User {
	var err error
	var user model.User

	if err = userDao.UserCollection.FindOne(util.GetHelper().GetContent(), bson.D{{"name", name}}).Decode(&user); err != nil {
		// TODO 错误处理
	}
	return user
}

func (userDao UserDao) UpdateUserByName(oldName, newName string) *mongo.UpdateResult {
	var err error
	var result *mongo.UpdateResult
	if result, err = userDao.UserCollection.UpdateOne(util.GetHelper().GetContent(), bson.M{"name": oldName}, bson.M{"$set": bson.M{"name": newName}}); err != nil {
		// TODO 错误处理
	}
	return result
}

func (userDao UserDao) RemoveUserByName(name string) *mongo.DeleteResult {
	var err error
	var result *mongo.DeleteResult
	if result, err = userDao.UserCollection.DeleteOne(util.GetHelper().GetContent(), bson.M{"name": name}); err != nil {
		// TODO 错误处理
	}
	return result
}
