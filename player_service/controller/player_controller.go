package controller

import (
	"encoding/json"
	"net/http"

	"github.com/berpergian/chi_learning/player_service/message"
	"github.com/berpergian/chi_learning/player_service/service"
	"github.com/berpergian/chi_learning/shared/config"
)

type PlayerController struct {
	PlayerService *service.PlayerService
	Env           *config.Env
}

func RegisterPlayerController(env *config.Env, playerService *service.PlayerService) *PlayerController {
	return &PlayerController{PlayerService: playerService, Env: env}
}

// GetList godoc
// @Summary      Player GetList
// @Description  Returns list of players
// @Tags         player
// @Accept       json
// @Produce      json
// @Param        pageSkip  query int false "Page skip (default: 0)"
// @Param        pageSize  query int false "Page size (default: 10)"
// @Success      200  {array} model.Player
// @Failure      500  {object} map[string]string
// @Security     BearerAuth
// @Router       /players [get]
func (controller *PlayerController) GetList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pageSkip := 1
	pageSize := 10

	players, err := controller.PlayerService.GetAllData(ctx, pageSkip, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message.ToPlayerGetListResponse(players))
}
