package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/berpergian/chi_learning/shared/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadDatabase(env *Env) database.IDatabaseClient {
	if true {
		return SetupMongoDatabase(env)
	}

	return SetupMongoDatabase(env)
}

// MongoDB - Start
func SetupMongoDatabase(env *Env) database.IDatabaseClient {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongodbURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
		CloseMongoDBConnection(*client)
	}

	err = client.Ping(ctx, opts.ReadPreference)
	if err != nil {
		log.Fatal(err)
	}

	return &database.MongoClient{Client: client}
}

func CloseMongoDBConnection(client mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}

// MongoDB - End
