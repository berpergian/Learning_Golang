package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPlayer = "players"
)

type Player struct {
	ID       primitive.ObjectID `bson:"_id"`
	PlayerId string             `bson:"playerId"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	BaseDoc  `bson:",inline"`
}

type PlayerModel interface {
	Create(c context.Context, player *Player) error
	Fetch(c context.Context) ([]Player, error)
	GetByEmail(c context.Context, email string) (Player, error)
	GetByID(c context.Context, id string) (Player, error)
	GetByPlayerID(c context.Context, playerId string) (Player, error)
}
