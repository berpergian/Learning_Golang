package service

import (
	"context"

	"github.com/berpergian/chi_learning/config"
	"github.com/berpergian/chi_learning/model"
	"github.com/berpergian/chi_learning/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlayerService struct {
	PlayerRepository *repository.PlayerRepository
	Env              *config.Env
}

func RegisterPlayerService(env *config.Env, playerRepository *repository.PlayerRepository) *PlayerService {
	return &PlayerService{PlayerRepository: playerRepository, Env: env}
}

func (service *PlayerService) GetAllData(ctx context.Context, pageSkip int, pageSize int) ([]model.Player, error) {
	if pageSkip == 0 {
		pageSkip = 1
	}

	if pageSize == 0 {
		pageSize = 10
	}

	limit := int64(pageSize)
	skip := int64((pageSkip - 1) * int(pageSize))

	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)

	players, err := service.PlayerRepository.Fetch(ctx, bson.M{}, *findOptions)
	if err != nil {
		return make([]model.Player, 1), err
	}

	return players, nil
}
