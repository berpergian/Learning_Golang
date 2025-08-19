package route

import (
	"time"

	"github.com/berpergian/chi_learning/config"
	"github.com/berpergian/chi_learning/controller"
	"github.com/berpergian/chi_learning/database"
	"github.com/berpergian/chi_learning/repository"
	"github.com/berpergian/chi_learning/service"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(timeout time.Duration, route *chi.Mux, env *config.Env, dbClient database.IDatabaseClient) {
	jwtManager := &service.JWTManager{
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
		AccountRouter(router, env, jwtManager, database)
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

func AccountRouter(router chi.Router, env *config.Env, jwtManager *service.JWTManager, database *mongo.Database) {
	playerRepository := repository.RegisterPlayerRepository(database)
	accountService := service.RegisterAccountService(env, playerRepository, jwtManager)
	accountController := controller.RegisterAccountController(env, accountService)

	router.Post("/registerOrLogin", accountController.RegisterOrLogin)
}

func PlayerRouter(router chi.Router, env *config.Env, database *mongo.Database) {
	playerRepository := repository.RegisterPlayerRepository(database)
	playerService := service.RegisterPlayerService(env, playerRepository)
	playerController := controller.RegisterPlayerController(env, playerService)

	router.Get("/players", playerController.GetList)
}
