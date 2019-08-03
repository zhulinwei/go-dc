package dao

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoDB struct {
	client *mongo.Client
}

//func GetContext() context.Context {
//	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
//	return ctx
//}

func (database *MongoDB) InitMongoDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	err = client.Connect(context.Background())
	if err != nil {
		panic(errors.New("mongodb connect fail"))
	}
	pingCtx, c1 := context.WithTimeout(context.Background(), 3*time.Second)
	defer c1()
	err = client.Ping(pingCtx, readpref.Primary())
	if err != nil {
		panic(errors.New("mongodb ping fail"))
	}
	database.client = client
}
