package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDB struct {
	client *mongo.Client
}

func (database *MongoDB) InitMongoDB() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))

	//fmt.Println(mongo.NewClient)
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	fmt.Println(client, err)
	//database.client = mongo
	//database.mongodb = mongo.
	//fmt.Println(database.mongodb)
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	//fmt.Println(client, err)
}
