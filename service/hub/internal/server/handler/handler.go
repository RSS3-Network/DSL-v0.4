package handler

import (
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

const (
	DefaultLimit = 500
)

type Handler struct {
	DatabaseClient     *gorm.DB
	RabbitmqConnection *rabbitmq.Connection
	RabbitmqChannel    *rabbitmq.Channel
}

type Response struct {
	Total  int64  `json:"total"`
	Cursor string `json:"cursor"`
	Result any    `json:"result"`
}
