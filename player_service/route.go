package main

import (
	"time"

	"github.com/berpergian/chi_learning/player_service/controller"
	"github.com/berpergian/chi_learning/player_service/service"
	"github.com/berpergian/chi_learning/shared/config"
	sharedService "github.com/berpergian/chi_learning/shared/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func RouteSetup(timeout time.Duration, route *chi.Mux, env *config.Env,
	playerService *service.PlayerService, bus *config.RabbitBus, validate *validator.Validate) {
	jwtManager := &sharedService.JWTManager{
		Secret: []byte(env.AccessTokenSecret),
		Issuer: env.Issuer,
		Expiry: time.Duration(env.AccessTokenExpiryHour) * time.Hour,
	}

	// Public APIs
	route.Group(func(router chi.Router) {
		HealthRouter(router)
	})

	// Secured APIs (Player)
	route.Group(func(router chi.Router) {
		router.Use(jwtManager.Middleware)

		PlayerRouter(router, env, playerService, validate)
	})

	// Secured APIs (Admin)
	route.Group(func(router chi.Router) {
		// TODO Change check jwt
		router.Use(jwtManager.Middleware)

		router.Route("/admin", func(subRoute chi.Router) {
			PlayerAdminRouter(subRoute, env, playerService, validate)
		})
	})
}

func HealthRouter(router chi.Router) {
	controller := controller.HealthController{}

	router.Get("/health", controller.Check)
}

func PlayerRouter(router chi.Router, env *config.Env, playerService *service.PlayerService, validate *validator.Validate) {
	controller := controller.RegisterPlayerController(env, playerService, validate)

	router.Get("/player", controller.GetInfo)
	router.Get("/player/inventories", controller.GetInventories)
	router.Get("/player/characters", controller.GetCharacterList)
}

func PlayerAdminRouter(router chi.Router, env *config.Env, playerService *service.PlayerService, validate *validator.Validate) {
	controller := controller.RegisterPlayerAdminController(env, playerService, validate)

	router.Get("/players", controller.GetList)
}
