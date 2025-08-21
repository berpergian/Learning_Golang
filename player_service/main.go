package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/berpergian/chi_learning/player_service/docs" // swagger generated docs
	"github.com/berpergian/chi_learning/player_service/subscriber"
	"github.com/berpergian/chi_learning/shared/config"
	sharedConstant "github.com/berpergian/chi_learning/shared/constant"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Player Service Learning API
// @version 1.0
// @description This is a sample description.
// @BasePath /

func main() {
	app := &config.Application{}
	app.Env = config.ReadEnvironment(sharedConstant.Player)
	app.Database = config.ReadDatabase(app.Env)

	bus, err := config.RegisterRabbitBus(app.Env)
	if err != nil {
		panic(err)
	}
	subscriber.Register(bus)
	defer bus.Close()
	defer app.Database.CloseDatabase()

	timeout := time.Duration(app.Env.ContextTimeout) * time.Second

	validate := validator.New(validator.WithRequiredStructEnabled())
	router := chi.NewRouter()

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

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost"+app.Env.ServerAddress+"/swagger/doc.json"),
	))

	RouteSetup(timeout, router, app.Env, app.Database, bus, validate)

	fmt.Println("Server up with environment:" + app.Env.AppEnv)
	http.ListenAndServe(app.Env.ServerAddress, router)
}
