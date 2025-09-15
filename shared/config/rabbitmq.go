package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitBus struct {
	Connection  *amqp.Connection
	Channel     *amqp.Channel
	Exchange    string
	ServiceName string
}

func RegisterRabbitBus(env *Env, serviceName string) (*RabbitBus, error) {
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
	if err := ch.ExchangeDeclare(exchange, "topic", true, false, false, false, nil); err != nil {
		return nil, err
	}
	return &RabbitBus{Connection: conn, Channel: ch, Exchange: exchange, ServiceName: serviceName}, nil
}

func (b *RabbitBus) Publish(ctx context.Context, eventName string, payload any) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	log.Printf("[" + b.ServiceName + "] Publish: " + eventName)
	return b.Channel.PublishWithContext(ctx, b.Exchange, "", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
		Type:        eventName,
	})
}

func (b *RabbitBus) Subscribe(queueName string, handler func(eventType string, body []byte) error) error {
	q, err := b.Channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	if err := b.Channel.QueueBind(q.Name, "", b.Exchange, false, nil); err != nil {
		return err
	}
	msgs, err := b.Channel.Consume(q.Name, "", true, false, false, false, nil)
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
	if b.Channel != nil {
		if err := b.Channel.Close(); err != nil {
			fmt.Println("channel close:", err)
		}
	}
	if b.Connection != nil {
		return b.Connection.Close()
	}
	return nil
}
