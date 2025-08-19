# Event-Driven Microservices with RabbitMQ

Services:
- account_service :8081 (POST /registerOrLogin)
- player_service  :8082 (GET /players)
- RabbitMQ: 5672 (AMQP), 15672 (UI)

Quick start:
1) docker run -d --name rabbit -p 5672:5672 -p 15672:15672 rabbitmq:3-management
2) cd account_service && go run ./...
3) cd player_service && go run ./...
