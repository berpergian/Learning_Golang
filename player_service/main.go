package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/berpergian/chi_learning/player_service/docs" // swagger generated docs
	"github.com/berpergian/chi_learning/player_service/repository"
	"github.com/berpergian/chi_learning/player_service/service"
	"github.com/berpergian/chi_learning/player_service/subscriber"
	"github.com/berpergian/chi_learning/shared/config"
	sharedConstant "github.com/berpergian/chi_learning/shared/constant"
	"github.com/berpergian/chi_learning/shared/database"
	"github.com/berpergian/chi_learning/shared/event"
	sharedStaticData "github.com/berpergian/chi_learning/shared/staticdata"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Player Service Learning API
// @version 1.0
// @description This is a sample description.
// @BasePath /

const Service string = "Player Service"

func main() {
	app := &config.Application{}
	app.Env = config.ReadEnvironment(sharedConstant.Player)
	app.Database = config.ReadDatabase(app.Env)

	staticDataService := sharedStaticData.InitializeStaticData()
	if err := staticDataService.Load("../shared/staticdata"); err != nil {
		panic(err)
	}

	database := database.GetMongoDatabase(app.Database, app.Env.DBName)
	playerRepository := repository.RegisterPlayerRepository(database)
	playerCharacterRepository := repository.RegisterPlayerCharacterRepository(database)
	playerInventoryRepository := repository.RegisterPlayerInventoryRepository(database)
	playerService := service.RegisterPlayerService(app.Env, playerRepository, playerCharacterRepository, playerInventoryRepository, staticDataService)

	bus, err := config.RegisterRabbitBus(app.Env, event.PlayerService)
	if err != nil {
		panic(err)
	}
	subscriber.Register(app.Env, bus, playerService)

	defer bus.Close()
	defer app.Database.CloseDatabase()

	timeout := time.Duration(app.Env.ContextTimeout) * time.Second

	validate := validator.New(validator.WithRequiredStructEnabled())
	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Timeout request
	router.Use(middleware.Timeout(60 * time.Second))

	// CORS middleware
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"}, // Use your allowed origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Origin", "X-Requested-With"},
		ExposedHeaders:   []string{"Link", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(cors.Handler)

	url := "http://localhost" + app.Env.ServerAddress
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(url+"/swagger/doc.json"),
	))

	RouteSetup(timeout, router, app.Env, playerService, bus, validate)

	fmt.Println("Running '" + Service + "' URL: " + url)
	fmt.Println("Server up with environment:" + app.Env.AppEnv)
	http.ListenAndServe(app.Env.ServerAddress, router)
}
