package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlayerBase struct {
	BaseModel `bson:",inline"`
	PlayerId  string `bson:"playerId"`
}

func CreatePlayerBase(id primitive.ObjectID, playerId string) PlayerBase {
	return PlayerBase{
		BaseModel: BaseModel{
			ID:        id,
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		PlayerId: playerId,
	}
}
