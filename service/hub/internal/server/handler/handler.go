package handler

import (
	"github.com/naturalselectionlabs/pregod/common/database"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Handler struct {
	DatabaseClient     *database.Client
	RabbitmqConnection *rabbitmq.Connection
	RabbitmqChannel    *rabbitmq.Channel
	TracerProvider     *trace.TracerProvider
}
