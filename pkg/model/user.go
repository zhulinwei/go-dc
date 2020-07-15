package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDB struct {
	ID   primitive.ObjectID `bson:"_id"`
	Age  int64              `bson:"age"`
	Name string             `bson:"name"`
}

type UserRequest struct {
	Age  int    `json:"age" binding:"required,min=0"`
	Name string `json:"name" binding:"required"`
}

type UserResponse struct {
	ID   primitive.ObjectID `json:"_id"`
	Age  int                `json:"age"`
	Name string             `json:"name"`
}
