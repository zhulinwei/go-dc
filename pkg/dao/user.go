package dao

import (
	"github.com/zhulinwei/go-dc/pkg/database"
	"github.com/zhulinwei/go-dc/pkg/model"
	"github.com/zhulinwei/go-dc/pkg/util"
	"github.com/zhulinwei/go-dc/pkg/util/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IUserDao interface {
	QueryUserByName(name string) model.UserDB
	SaveUser(user model.UserRequest) *mongo.InsertOneResult
	BulkSaveUser(users []model.UserRequest) *mongo.InsertManyResult
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

func (userDao UserDao) SaveUser(user model.UserRequest) *mongo.InsertOneResult {
	var err error
	var result *mongo.InsertOneResult
	if result, err = userDao.UserCollection.InsertOne(util.CommonContent(), user); err != nil {
		log.Error("userDao save user fail", log.String("error", err.Error()))
	}
	log.Debug("save user success", log.Reflect("result", result))
	return result
}


func (userDao UserDao) BulkSaveUser(users []model.UserRequest) *mongo.InsertManyResult {
	//var err error
	var result *mongo.InsertManyResult
	// insert documents {name: "Alice"} and {name: "Bob"}
	// set the Ordered option to false to allow both operations to happen even if one of them errors
	//docs := []interface{}{
	//	bson.D{{"name", "Alice"}},
	//	bson.D{{"name", "Bob"}},
	//}
	// SetOrdered为false指的保存失败的数据不会影响正常保存的数据
	//opts := options.InsertMany().SetOrdered(false)
	//if result, err = userDao.UserCollection.InsertMany(util.CommonContent(), []interface{}{users}, opts); err != nil {
	//	log.Error("userDao bulk save user fail", log.String("error", err.Error()))
	//}

	var models []mongo.WriteModel
	for i := 0; i < len(users); i++ {
		modelx := mongo.NewInsertOneModel()
		modelx.SetDocument(users[i])
		models = append(models, modelx)
	}
	opts2 := options.BulkWrite().SetOrdered(false)
	_, _ = userDao.UserCollection.BulkWrite(util.CommonContent(), models, opts2)

	return result
}


func (userDao UserDao) QueryUserByName(name string) model.UserDB {
	var err error
	var user model.UserDB

	if err = userDao.UserCollection.FindOne(util.CommonContent(), bson.D{{"name", name}}).Decode(&user); err != nil {
		log.Error("userDao find user fail", log.String("error", err.Error()))
	}
	return user
}

func (userDao UserDao) UpdateUserByName(oldName, newName string) *mongo.UpdateResult {
	var err error
	var result *mongo.UpdateResult
	if result, err = userDao.UserCollection.UpdateOne(util.CommonContent(), bson.M{"name": oldName}, bson.M{"$set": bson.M{"name": newName}}); err != nil {
		log.Error("userDao update user fail", log.String("error", err.Error()))
	}
	return result
}

func (userDao UserDao) RemoveUserByName(name string) *mongo.DeleteResult {
	var err error
	var result *mongo.DeleteResult
	if result, err = userDao.UserCollection.DeleteOne(util.CommonContent(), bson.M{"name": name}); err != nil {
		log.Error("userDao delete user fail", log.String("error", err.Error()))
	}
	return result
}
