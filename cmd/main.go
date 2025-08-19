package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/berpergian/chi_learning/config"
	"github.com/berpergian/chi_learning/database"
	_ "github.com/berpergian/chi_learning/docs" // swagger generated docs
	"github.com/berpergian/chi_learning/route"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Chi Learning API
// @version 1.0
// @description This is a sample description.
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

type Application struct {
	Env      *config.Env
	Database database.IDatabaseClient
}

func main() {
	app := &Application{}
	app.Env = config.ReadEnvironment()
	app.Database = config.ReadDatabase(app.Env)

	defer app.Database.CloseDatabase()

	timeout := time.Duration(app.Env.ContextTimeout) * time.Second

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

	route.Setup(timeout, router, app.Env, app.Database)

	fmt.Println("Server up with environment:" + app.Env.AppEnv)
	http.ListenAndServe(app.Env.ServerAddress, router)
}
