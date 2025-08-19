package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClient struct {
	Client *mongo.Client
}

func (mongoClient *MongoClient) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return mongoClient.Client.Ping(ctx, nil)
}

func (mongoClient *MongoClient) CloseDatabase() {
	err := mongoClient.Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}

func (mongoClient *MongoClient) GetDatabase(databaseName string) *mongo.Database {
	return mongoClient.Client.Database(databaseName)
}
