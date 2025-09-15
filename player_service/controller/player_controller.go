package controller

import (
	"encoding/json"
	"net/http"

	"github.com/berpergian/chi_learning/player_service/service"
	"github.com/berpergian/chi_learning/shared/config"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"

	sharedHelper "github.com/berpergian/chi_learning/shared/helper"
	sharedService "github.com/berpergian/chi_learning/shared/service"
)

type PlayerController struct {
	PlayerService *service.PlayerService
	Env           *config.Env
	Validate      *validator.Validate
}

func RegisterPlayerController(env *config.Env, playerService *service.PlayerService, validate *validator.Validate) *PlayerController {
	return &PlayerController{PlayerService: playerService, Env: env, Validate: validate}
}

// GetInfo godoc
// @Summary      Player GetInfo
// @Description  Returns general info of player
// @Tags         player
// @Accept       json
// @Produce      json
// @Success      200  {object} message.PlayerInfoResponse
// @Failure      500  {object} map[string]string
// @Security     BearerAuth
// @Router       /players [get]
func (controller *PlayerController) GetInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	playerId, isExist := sharedService.TryGetPlayerIDFromContext(ctx)
	if !isExist {
		sharedHelper.WriteProblem(w, r, sharedHelper.ProblemDetails{
			Title:     "Unauthorized",
			Status:    http.StatusUnauthorized,
			Detail:    "Cannot get player data",
			RequestID: r.Context().Value(middleware.RequestIDKey),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playerId)
}

// GetInventories godoc
// @Summary      Player GetInventories
// @Description  Returns inventory player
// @Tags         player
// @Accept       json
// @Produce      json
// @Success      200  {object} message.PlayerInfoResponse
// @Failure      500  {object} map[string]string
// @Security     BearerAuth
// @Router       /players [get]
func (controller *PlayerController) GetInventories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	playerId, isExist := sharedService.TryGetPlayerIDFromContext(ctx)
	if !isExist {
		sharedHelper.WriteProblem(w, r, sharedHelper.ProblemDetails{
			Title:     "Unauthorized",
			Status:    http.StatusUnauthorized,
			Detail:    "Cannot get player data",
			RequestID: r.Context().Value(middleware.RequestIDKey),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playerId)
}

// GetCharacterList godoc
// @Summary      Player GetCharacterList
// @Description  Returns character list player
// @Tags         player
// @Accept       json
// @Produce      json
// @Success      200  {object} message.PlayerInfoResponse
// @Failure      500  {object} map[string]string
// @Security     BearerAuth
// @Router       /players [get]
func (controller *PlayerController) GetCharacterList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	playerId, isExist := sharedService.TryGetPlayerIDFromContext(ctx)
	if !isExist {
		sharedHelper.WriteProblem(w, r, sharedHelper.ProblemDetails{
			Title:     "Unauthorized",
			Status:    http.StatusUnauthorized,
			Detail:    "Cannot get player data",
			RequestID: r.Context().Value(middleware.RequestIDKey),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playerId)
}
