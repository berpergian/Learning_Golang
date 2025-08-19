package service

import (
	"context"
	"errors"

	"github.com/berpergian/chi_learning/config"
	"github.com/berpergian/chi_learning/message"
	"github.com/berpergian/chi_learning/model"
	"github.com/berpergian/chi_learning/repository"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AccountService struct {
	PlayerRepository *repository.PlayerRepository
	JWTManager       *JWTManager
	Env              *config.Env
}

func RegisterAccountService(env *config.Env, playerRepository *repository.PlayerRepository, jwtManager *JWTManager) *AccountService {
	return &AccountService{PlayerRepository: playerRepository, JWTManager: jwtManager, Env: env}
}

func (s *AccountService) RegisterOrLogin(ctx context.Context, req message.LoginRequest) (*message.LoginResponse, error) {
	player, err := s.PlayerRepository.GetByEmail(ctx, req.Email)
	if err != nil || player.ID.IsZero() {
		// Register new player
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		player = model.Player{
			ID:       primitive.NewObjectID(),
			PlayerId: uuid.New().String(),
			Name:     req.Name,
			Email:    req.Email,
			Password: string(hashedPassword),
		}
		if err := s.PlayerRepository.Create(ctx, &player); err != nil {
			return nil, err
		}
	} else {
		// Existing player, check password
		if err := bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(req.Password)); err != nil {
			return nil, errors.New("invalid credentials")
		}
	}

	// Generate JWT access token
	token, err := s.JWTManager.Generate(player.PlayerId)
	if err != nil {
		return nil, err
	}

	return &message.LoginResponse{
		AccessToken: token,
		PlayerId:    player.PlayerId,
		Email:       player.Email,
		Name:        player.Name,
	}, nil
}
