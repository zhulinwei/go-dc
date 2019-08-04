package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Test1Model struct {
	Test1ID primitive.ObjectID `bson:"_id"`
	Age     int64              `bson:"age"`
	Name    string             `bson:"name`
}
