package service

import (
	"context"
	"log"
	"time"

	"github.com/berpergian/chi_learning/player_service/repository"
	"github.com/berpergian/chi_learning/shared/config"
	"github.com/berpergian/chi_learning/shared/event"
	"github.com/berpergian/chi_learning/shared/model"
	sharedStaticData "github.com/berpergian/chi_learning/shared/staticdata"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlayerService struct {
	Env                       *config.Env
	PlayerRepository          *repository.PlayerRepository
	PlayerCharacterRepository *repository.PlayerCharacterRepository
	PlayerInventoryRepository *repository.PlayerInventoryRepository
	StaticDataService         *sharedStaticData.StaticDataService
}

func RegisterPlayerService(env *config.Env,
	playerRepository *repository.PlayerRepository,
	playerCharacterRepository *repository.PlayerCharacterRepository,
	playerInventoryRepository *repository.PlayerInventoryRepository,
	staticDataService *sharedStaticData.StaticDataService) *PlayerService {
	return &PlayerService{
		Env:                       env,
		PlayerRepository:          playerRepository,
		PlayerCharacterRepository: playerCharacterRepository,
		PlayerInventoryRepository: playerInventoryRepository,
		StaticDataService:         staticDataService,
	}
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

func (service *PlayerService) SetupPlayerRegistered(data event.PlayerRegistered) {
	characters := service.StaticDataService.Characters().All()
	items := service.StaticDataService.Items().All()

	var charaInvent []model.CharacterStaticData
	for _, chara := range characters {
		chara.Level = 1
		charaInvent = append(charaInvent, chara)
	}

	playerChara := model.PlayerCharacter{
		PlayerBase: model.CreatePlayerBase(primitive.NewObjectID(), data.PlayerID),
		Items:      charaInvent,
	}
	playerInvent := model.PlayerInventory{
		PlayerBase: model.CreatePlayerBase(primitive.NewObjectID(), data.PlayerID),
		Items:      items,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if errChara := service.PlayerCharacterRepository.Create(ctx, &playerChara); errChara != nil {
		log.Println(errChara)
	}
	if errInvent := service.PlayerInventoryRepository.Create(ctx, &playerInvent); errInvent != nil {
		log.Println(errInvent)
	}
}
