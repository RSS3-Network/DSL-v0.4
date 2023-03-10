package config

import (
	configx "github.com/naturalselectionlabs/pregod/common/config"
)

type Config struct {
	Mode          configx.Mode           `mapstructure:"mode"`
	Postgres      *configx.Postgres      `mapstructure:"postgres"`
	RabbitMQ      *configx.RabbitMQ      `mapstructure:"rabbitmq"`
	OpenTelemetry *configx.OpenTelemetry `mapstructure:"opentelemetry"`
	Redis         *configx.Redis         `mapstructure:"redis"`
	RPC           *configx.RPC           `mapstructure:"rpc"`
	Kurora        *configx.Kurora        `mapstructure:"kurora"`
	EIP1577       *configx.EIP1577       `mapstructure:"eip1577"`
	Sound         *configx.Sound         `mapstructure:"sound"`
}
