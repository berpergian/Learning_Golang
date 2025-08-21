package main

import (
	"time"

	"github.com/berpergian/chi_learning/account_service/controller"
	"github.com/berpergian/chi_learning/account_service/repository"
	"github.com/berpergian/chi_learning/account_service/service"
	"github.com/berpergian/chi_learning/shared/config"
	"github.com/berpergian/chi_learning/shared/database"
	sharedService "github.com/berpergian/chi_learning/shared/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

func RouteSetup(timeout time.Duration, route *chi.Mux, env *config.Env,
	dbClient database.IDatabaseClient, bus *config.RabbitBus, validate *validator.Validate) {
	jwtManager := &sharedService.JWTManager{
		Secret: []byte(env.AccessTokenSecret),
		Issuer: env.Issuer,
		Expiry: time.Duration(env.AccessTokenExpiryHour) * time.Hour,
	}

	mongoClient, ok := dbClient.(*database.MongoClient)
	if !ok {
		panic("unsupported database client type")
	}
	database := mongoClient.Client.Database(env.DBName)

	// Public APIs
	route.Group(func(router chi.Router) {
		HealthRouter(router)
		AccountRouter(router, env, jwtManager, database, bus, validate)
	})
}

func HealthRouter(router chi.Router) {
	healthController := controller.HealthController{}

	router.Get("/health", healthController.Check)
}

func AccountRouter(router chi.Router, env *config.Env, jwtManager *sharedService.JWTManager,
	database *mongo.Database, bus *config.RabbitBus, validate *validator.Validate) {
	playerRepository := repository.RegisterPlayerRepository(database)
	accountService := service.RegisterAccountService(env, playerRepository, jwtManager, bus, validate)
	accountController := controller.RegisterAccountController(env, accountService)

	router.Post("/register", accountController.Register)
	router.Post("/login", accountController.Login)
}
