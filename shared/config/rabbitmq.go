package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitBus struct {
	conn     *amqp.Connection
	ch       *amqp.Channel
	exchange string
}

func RegisterRabbitBus(env *Env) (*RabbitBus, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s", env.RabbitUser, env.RabbitPass, env.RabbitHost, env.RabbitPort)
	exchange := env.MessageExchange

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	if err := ch.ExchangeDeclare(exchange, "fanout", true, false, false, false, nil); err != nil {
		return nil, err
	}
	return &RabbitBus{conn: conn, ch: ch, exchange: exchange}, nil
}

func (b *RabbitBus) Publish(ctx context.Context, eventName string, payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return b.ch.PublishWithContext(ctx, b.exchange, "", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
		Type:        eventName,
	})
}

func (b *RabbitBus) Subscribe(queueName string, handler func(eventType string, body []byte) error) error {
	q, err := b.ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	if err := b.ch.QueueBind(q.Name, "", b.exchange, false, nil); err != nil {
		return err
	}
	msgs, err := b.ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		for d := range msgs {
			if err := handler(d.Type, d.Body); err != nil {
				log.Printf("handler error: %v", err)
			}
		}
	}()
	return nil
}

func (b *RabbitBus) Close() error {
	if b.ch != nil {
		if err := b.ch.Close(); err != nil {
			fmt.Println("channel close:", err)
		}
	}
	if b.conn != nil {
		return b.conn.Close()
	}
	return nil
}
