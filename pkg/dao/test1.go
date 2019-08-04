package dao

import (
	"github.com/zhulinwei/gin-demo/pkg/dto"
	"github.com/zhulinwei/gin-demo/pkg/model"
	"github.com/zhulinwei/gin-demo/pkg/repository"
	"github.com/zhulinwei/gin-demo/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Test1Dao struct{}

var test1Dao = new(Test1Dao)

func GetTest1Dao() *Test1Dao {
	return test1Dao
}

func (*Test1Dao) SaveUser(test1 dto.User) *mongo.InsertOneResult {
	var err error
	var result *mongo.InsertOneResult
	if result, err = repository.GetMongoDBCollection().InsertOne(util.GetUitl().GetContent(), test1); err != nil {
		// TODO 错误处理
	}
	return result
}

func (*Test1Dao) QueryUserByName(name string) model.User {
	var err error
	var user model.User
	if err = repository.GetMongoDBCollection().FindOne(util.GetUitl().GetContent(), bson.D{{"name", name}}).Decode(&user); err != nil {
		// TODO 错误处理
	}
	return user
}

func (*Test1Dao) UpdateUserByName(oldName, newName string) *mongo.UpdateResult {
	var err error
	var result *mongo.UpdateResult
	if result, err = repository.GetMongoDBCollection().UpdateOne(util.GetUitl().GetContent(), bson.M{"name": oldName}, bson.M{"$set": bson.M{"name": newName}}); err != nil {
		// TODO 错误处理
	}
	return result
}

func (*Test1Dao) RemoveUserByName(name string) *mongo.DeleteResult {
	var err error
	var result *mongo.DeleteResult
	if result, err = repository.GetMongoDBCollection().DeleteOne(util.GetUitl().GetContent(), bson.M{"name": name}); err != nil {
		// TODO 错误处理
	}
	return result
}
