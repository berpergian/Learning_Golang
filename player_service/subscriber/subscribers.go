package subscriber

import (
	"encoding/json"
	"log"

	"github.com/berpergian/chi_learning/player_service/service"
	"github.com/berpergian/chi_learning/shared/config"
	"github.com/berpergian/chi_learning/shared/event"
)

func Register(env *config.Env, bus *config.RabbitBus, service *service.PlayerService) error {
	return bus.Subscribe(event.PlayerService, func(eventType string, body []byte) error {
		switch eventType {
		case event.PlayerRegisteredTopic:
			var e event.PlayerRegistered
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			log.Printf("["+event.PlayerService+"] PlayerRegistered: id=%s email=%s name=%s", e.PlayerID, e.Email, e.Name)
			service.SetupPlayerRegistered(e)
		case event.PlayerLoggedInTopic:
			var e event.PlayerLoggedIn
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			log.Printf("["+event.PlayerService+"] PlayerLoggedIn: id=%s email=%s", e.PlayerID, e.Email)
		}
		return nil
	})
}
