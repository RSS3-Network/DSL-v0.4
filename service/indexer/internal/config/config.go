package config

import (
	configx "github.com/naturalselectionlabs/pregod/common/config"
)

type Config struct {
	RabbitMQ      *configx.RabbitMQ      `mapstructure:"rabbitmq"`
	Postgres      *configx.Postgres      `mapstructure:"postgres"`
	OpenTelemetry *configx.OpenTelemetry `mapstructure:"opentelemetry"`
	Moralis       *configx.Moralis       `mapstructure:"moralis"`
	Redis         *configx.Redis         `mapstructure:"redis"`
	CoinMarketCap *configx.CoinMarketCap `mapstructure:"coinmarketcap"`
	Infura        *configx.Infura        `mapstructure:"infura"`
}
