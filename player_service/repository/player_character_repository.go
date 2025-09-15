package repository

import (
	sharedConstant "github.com/berpergian/chi_learning/shared/constant"
	"github.com/berpergian/chi_learning/shared/model"
	"github.com/berpergian/chi_learning/shared/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayerCharacterRepository struct {
	*repository.BaseRepository[model.PlayerCharacter]
}

func RegisterPlayerCharacterRepository(db *mongo.Database) *PlayerCharacterRepository {
	return &PlayerCharacterRepository{
		BaseRepository: repository.RegisterBaseRepository[model.PlayerCharacter](db, sharedConstant.CollectionPlayer, sharedConstant.PlayerCharacterDocument),
	}
}
