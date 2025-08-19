package repository

import (
	"context"

	"github.com/berpergian/chi_learning/player_service/model"
	"github.com/berpergian/chi_learning/shared/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlayerRepository struct {
	*repository.BaseRepository[model.Player]
}

func RegisterPlayerRepository(db *mongo.Database) *PlayerRepository {
	return &PlayerRepository{
		BaseRepository: repository.RegisterBaseRepository[model.Player](db, model.CollectionPlayer),
	}
}

func (repository *PlayerRepository) GetByEmail(ctx context.Context, email string) (model.Player, error) {
	return repository.GetOne(ctx, bson.M{"email": email})
}

func (repository *PlayerRepository) GetByPlayerID(ctx context.Context, playerId string) (model.Player, error) {
	return repository.GetOne(ctx, bson.M{"playerId": playerId})
}
