package repository

import (
	sharedConstant "github.com/berpergian/chi_learning/shared/constant"
	"github.com/berpergian/chi_learning/shared/model"
	"github.com/berpergian/chi_learning/shared/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayerInventoryRepository struct {
	*repository.BaseRepository[model.PlayerInventory]
}

func RegisterPlayerInventoryRepository(db *mongo.Database) *PlayerInventoryRepository {
	return &PlayerInventoryRepository{
		BaseRepository: repository.RegisterBaseRepository[model.PlayerInventory](db, sharedConstant.CollectionPlayer, sharedConstant.PlayerInventoryDocument),
	}
}
