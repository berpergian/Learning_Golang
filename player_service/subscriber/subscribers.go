package subscriber

import (
	"encoding/json"
	"log"

	"github.com/berpergian/chi_learning/shared/config"
	"github.com/berpergian/chi_learning/shared/event"
)

func Register(bus *config.RabbitBus) error {
	return bus.Subscribe("player-service", func(eventType string, body []byte) error {
		switch eventType {
		case event.TopicPlayerRegistered:
			var e event.PlayerRegistered
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			log.Printf("[player-service] PlayerRegistered: id=%s email=%s name=%s", e.PlayerID, e.Email, e.Name)
		case event.TopicPlayerLoggedIn:
			var e event.PlayerLoggedIn
			if err := json.Unmarshal(body, &e); err != nil {
				return err
			}
			log.Printf("[player-service] PlayerLoggedIn: id=%s email=%s", e.PlayerID, e.Email)
		}
		return nil
	})
}
