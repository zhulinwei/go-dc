package dao

import (
	"context"

	"github.com/zhulinwei/go-dc/pkg/database"
	"github.com/zhulinwei/go-dc/pkg/model"
	"github.com/zhulinwei/go-dc/pkg/util"
	"github.com/zhulinwei/go-dc/pkg/util/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IUserDao interface {
	SaveUser(user model.UserRequest) (*mongo.InsertOneResult, error)
	BulkSaveUser(users []model.UserRequest) (*mongo.BulkWriteResult, error)
	QueryUserByName(name string) (*model.UserDB, error)
	QueryUsersByName(name string) ([]model.UserDB, error)
	RemoveUserByName(name string) (*mongo.DeleteResult, error)
	UpdateUserByName(oldName, newName string) (*mongo.UpdateResult, error)
}

type UserDao struct {
	UserCollection *mongo.Collection
}

func BuildUserDao() IUserDao {
	return UserDao{
		UserCollection: database.BuildMongoDB().UserCollection(),
	}
}

func (userDao UserDao) SaveUser(user model.UserRequest) (*mongo.InsertOneResult, error) {
	result, err := userDao.UserCollection.InsertOne(util.CommonContent(), user)
	if err != nil {
		log.Error("userDao save user fail", log.String("error", err.Error()))
		return nil, err
	}
	log.Debug("save user success", log.Reflect("result", result))
	return result, nil
}

func (userDao UserDao) BulkSaveUser(users []model.UserRequest) (*mongo.BulkWriteResult, error) {
	var models []mongo.WriteModel
	for i := 0; i < len(users); i++ {
		models = append(models, mongo.NewInsertOneModel().SetDocument(users[i]))
	}
	opts := options.BulkWrite().SetOrdered(false)
	result, err := userDao.UserCollection.BulkWrite(util.CommonContent(), models, opts)
	if err != nil {
		log.Error("bulk save user fail", log.Reflect("error", err.Error()))
		return nil, err
	}
	return result, nil
}

func (userDao UserDao) QueryUserByName(name string) (*model.UserDB, error) {
	var user = new(model.UserDB)
	if err := userDao.UserCollection.FindOne(util.CommonContent(), bson.D{{"name", name}}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Error("userDao find user fail", log.String("error", err.Error()))
		return nil, err
	}
	return user, nil
}

func (userDao UserDao) QueryUsersByName(name string) ([]model.UserDB, error) {
	ops := &options.FindOptions{}
	cursor, err := userDao.UserCollection.Find(util.CommonContent(), bson.D{{"name", name}}, ops)

	if err != nil {
		log.Error("userDao find user fail", log.String("error", err.Error()))
		return nil, err
	}

	var users []model.UserDB
	for cursor.Next(context.TODO()) {
		var user model.UserDB
		if err := cursor.Decode(&user); err != nil {
			log.Error("userDao cursor error", log.String("error", err.Error()))
			return nil, nil
		}
		users = append(users, user)
	}
	return users, nil
}

func (userDao UserDao) UpdateUserByName(oldName, newName string) (*mongo.UpdateResult, error) {
	result, err := userDao.UserCollection.UpdateOne(util.CommonContent(), bson.M{"name": oldName}, bson.M{"$set": bson.M{"name": newName}})
	if err != nil {
		log.Error("userDao update user fail", log.String("error", err.Error()))
		return nil, err
	}
	return result, nil
}

func (userDao UserDao) RemoveUserByName(name string) (*mongo.DeleteResult, error) {
	result, err := userDao.UserCollection.DeleteOne(util.CommonContent(), bson.M{"name": name})
	if err != nil {
		log.Error("userDao delete user fail", log.String("error", err.Error()))
		return nil, err
	}
	return result, nil
}
