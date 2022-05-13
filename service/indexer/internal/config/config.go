package config

import "github.com/naturalselectionlabs/pregod/common/config"

type Config struct {
	RabbitMQ *configx.RabbitMQ `mapstructure:"rabbitmq"`
	Moralis  *configx.Moralis  `mapstructure:"moralis"`
}
