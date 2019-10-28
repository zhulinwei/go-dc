package dao

import (
	"github.com/zhulinwei/go-dc/pkg/database"
	"github.com/zhulinwei/go-dc/pkg/model"
	"github.com/zhulinwei/go-dc/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserDao interface {
	QueryUserByName(name string) model.UserDB
	SaveUser(test1 model.UserReq) *mongo.InsertOneResult
	RemoveUserByName(name string) *mongo.DeleteResult
	UpdateUserByName(oldName, newName string) *mongo.UpdateResult
}

type UserDao struct {
	UserCollection *mongo.Collection
}

func BuildUserDao() IUserDao {
	mongodb := database.BuildMongoDB()
	return UserDao{
		UserCollection: mongodb.UserCollection(),
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

func (userDao UserDao) QueryUserByName(name string) model.UserDB {
	var err error
	var user model.UserDB

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
