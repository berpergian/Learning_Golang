package message

import (
	"github.com/berpergian/chi_learning/model"
)

type PlayerGetListResponse struct {
	PlayerId string `json:"playerId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func ToPlayerGetListResponse(players []model.Player) []PlayerGetListResponse {
	responses := make([]PlayerGetListResponse, len(players))
	for i, p := range players {
		responses[i] = PlayerGetListResponse{
			PlayerId: p.PlayerId,
			Name:     p.Name,
			Email:    p.Email,
		}
	}
	return responses
}
