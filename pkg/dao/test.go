package dao

import (
	"github.com/zhulinwei/gin-demo/pkg/dto"
	"github.com/zhulinwei/gin-demo/pkg/model"
	"github.com/zhulinwei/gin-demo/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ITestDao interface {
	QueryUserByName(name string) model.Test
	SaveUser(test1 dto.Test) *mongo.InsertOneResult
	RemoveUserByName(name string) *mongo.DeleteResult
	UpdateUserByName(oldName, newName string) *mongo.UpdateResult
}

type TestDao struct {
	TestCollection *mongo.Collection
}

func NewTestDao (mongodb *mongo.Client) *TestDao {
	return &TestDao{
		TestCollection: mongodb.Database("test_database").Collection("test_collection"),
	}
}

func (testDao *TestDao) SaveUser(test1 dto.Test) *mongo.InsertOneResult {
	var err error
	var result *mongo.InsertOneResult
	if result, err = testDao.TestCollection.InsertOne(util.GetHelper().GetContent(), test1); err != nil {
		// TODO 错误处理
	}
	return result
}

func (testDao *TestDao) QueryUserByName(name string) model.Test {
	var err error
	var test model.Test

	if err = testDao.TestCollection.FindOne(util.GetHelper().GetContent(), bson.D{{"name", name}}).Decode(&test); err != nil {
		// TODO 错误处理
	}
	return test
}

func (testDao *TestDao) UpdateUserByName(oldName, newName string) *mongo.UpdateResult {
	var err error
	var result *mongo.UpdateResult
	if result, err = testDao.TestCollection.UpdateOne(util.GetHelper().GetContent(), bson.M{"name": oldName}, bson.M{"$set": bson.M{"name": newName}}); err != nil {
		// TODO 错误处理
	}
	return result
}

func (testDao *TestDao) RemoveUserByName(name string) *mongo.DeleteResult {
	var err error
	var result *mongo.DeleteResult
	if result, err = testDao.TestCollection.DeleteOne(util.GetHelper().GetContent(), bson.M{"name": name}); err != nil {
		// TODO 错误处理
	}
	return result
}
