package config

import "github.com/naturalselectionlabs/pregod/common/config"

type Config struct {
	HTTP          *configx.HTTP
	RabbitMQ      *configx.RabbitMQ
	OpenTelemetry *configx.OpenTelemetry
	Postgres      *configx.Postgres
}
