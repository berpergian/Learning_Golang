package model

import "go.mongodb.org/mongo-driver/bson/primitive"

const CollectionPlayer = "players"

type Player struct {
	ID       primitive.ObjectID `bson:"_id"`
	PlayerId string             `bson:"playerId"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}
