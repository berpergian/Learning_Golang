package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMongoDatabase(dbClient IDatabaseClient, dbName string) *mongo.Database {
	mongoClient, ok := dbClient.(*MongoClient)
	if !ok {
		panic("unsupported database client type")
	}
	return mongoClient.Client.Database(dbName)
}
