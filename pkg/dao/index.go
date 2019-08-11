package dao

import "go.mongodb.org/mongo-driver/mongo"

func InitDao (mongodb *mongo.Client) IUserDao {
	return NewUserDao(mongodb)
}
