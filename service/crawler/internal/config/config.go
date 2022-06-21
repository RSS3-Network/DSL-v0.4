package config

import (
	configx "github.com/naturalselectionlabs/pregod/common/config"
)

type Config struct {
	Postgres      *configx.Postgres      `mapstructure:"postgres"`
	RabbitMQ      *configx.RabbitMQ      `mapstructure:"rabbitmq"`
	OpenTelemetry *configx.OpenTelemetry `mapstructure:"opentelemetry"`
	Redis         *configx.Redis         `mapstructure:"redis"`
	Gateway       *configx.Gateway       `mapstructure:"gateway"`
}
