package message

import "github.com/berpergian/chi_learning/player_service/model"

type PlayerGetListResponse struct {
	PlayerId string `json:"playerId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func ToPlayerGetListResponse(players []model.Player) []PlayerGetListResponse {
	res := make([]PlayerGetListResponse, len(players))
	for i, p := range players {
		res[i] = PlayerGetListResponse{PlayerId: p.PlayerId, Name: p.Name, Email: p.Email}
	}
	return res
}
