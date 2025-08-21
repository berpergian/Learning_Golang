package main

import (
	"time"

	"github.com/berpergian/chi_learning/player_service/controller"
	"github.com/berpergian/chi_learning/player_service/repository"
	"github.com/berpergian/chi_learning/player_service/service"
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
	})

	// Secured APIs
	route.Group(func(router chi.Router) {
		router.Use(jwtManager.Middleware)

		PlayerRouter(router, env, database)
	})
}

func HealthRouter(router chi.Router) {
	healthController := controller.HealthController{}

	router.Get("/health", healthController.Check)
}

func PlayerRouter(router chi.Router, env *config.Env, database *mongo.Database) {
	playerRepository := repository.RegisterPlayerRepository(database)
	playerService := service.RegisterPlayerService(env, playerRepository)
	playerController := controller.RegisterPlayerController(env, playerService)

	router.Get("/players", playerController.GetList)
}
