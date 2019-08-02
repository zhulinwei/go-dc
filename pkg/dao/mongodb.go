package database

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	client *mongo.Client
}

func (database *MongoDB) InitMongoDB() {
	client, err := NewClient(options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	fmt.Println(client, err)
	//database.client = mongo
	//database.mongodb = mongo.
	//fmt.Println(database.mongodb)
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	//fmt.Println(client, err)
}