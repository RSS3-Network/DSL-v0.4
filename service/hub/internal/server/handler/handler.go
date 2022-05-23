package handler

import (
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type Handler struct {
	DatabaseClient     *gorm.DB
	RabbitmqConnection *rabbitmq.Connection
	RabbitmqChannel    *rabbitmq.Channel
}
