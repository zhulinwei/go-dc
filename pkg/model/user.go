package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDB struct {
	Test1ID primitive.ObjectID `bson:"_id"`
	Age     int64              `bson:"age"`
	Name    string             `bson:"name"`
}

type UserReq struct {
	Age  int    `json:"age" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UserRes struct {
	Test1ID primitive.ObjectID `json:"_id"`
	Age     int64              `json:"age"`
	Name    string             `json:"name"`
}
