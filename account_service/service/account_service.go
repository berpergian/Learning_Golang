package service

import (
	"context"
	"errors"

	"github.com/berpergian/chi_learning/account_service/message"
	"github.com/berpergian/chi_learning/account_service/repository"
	"github.com/berpergian/chi_learning/shared/config"
	"github.com/berpergian/chi_learning/shared/event"
	"github.com/berpergian/chi_learning/shared/model"
	"github.com/berpergian/chi_learning/shared/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AccountService struct {
	Repo     *repository.PlayerRepository
	JWT      *service.JWTManager
	Bus      *config.RabbitBus
	Env      *config.Env
	Validate *validator.Validate
}

func RegisterAccountService(env *config.Env, repo *repository.PlayerRepository,
	jwt *service.JWTManager, bus *config.RabbitBus, validate *validator.Validate) *AccountService {
	return &AccountService{Repo: repo, JWT: jwt, Bus: bus, Env: env, Validate: validate}
}

func (service *AccountService) Register(ctx context.Context, req message.RegisterRequest) (*message.RegisterResponse, error) {
	err := service.Validate.Struct(req)
	if err != nil {
		return nil, errors.New("request not valid")
	}

	player, err := service.Repo.GetByEmail(ctx, req.Email)
	if err != nil || player.ID.IsZero() {
		hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		player = model.Player{
			ID:       primitive.NewObjectID(),
			PlayerId: uuid.New().String(),
			Name:     req.Name,
			Email:    req.Email,
			Password: string(hashed),
		}

		if err := service.Repo.Create(ctx, &player); err != nil {
			return nil, err
		}

		_ = service.Bus.Publish(ctx, event.TopicPlayerRegistered, event.PlayerRegistered{
			PlayerID: player.PlayerId, Email: player.Email, Name: player.Name,
		})

		token, err := service.JWT.Generate(player.PlayerId)
		if err != nil {
			return nil, err
		}

		return &message.RegisterResponse{
			AccessToken: token, PlayerId: player.PlayerId, Email: player.Email, Name: player.Name,
		}, nil
	}

	return nil, errors.New("player already registered")
}

func (service *AccountService) Login(ctx context.Context, req message.LoginRequest) (*message.LoginResponse, error) {
	err := service.Validate.Struct(req)
	if err != nil {
		return nil, errors.New("request not valid")
	}

	player, err := service.Repo.GetByEmail(ctx, req.Email)
	if err != nil || player.ID.IsZero() {
		return nil, errors.New("player not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("wrong password")
	}

	token, err := service.JWT.Generate(player.PlayerId)
	if err != nil {
		return nil, err
	}

	return &message.LoginResponse{
		AccessToken: token,
	}, nil
}
