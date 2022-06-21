package config

import (
	configx "github.com/naturalselectionlabs/pregod/common/config"
)

type Config struct {
	HTTP          *configx.HTTP          `mapstructure:"http"`
	RabbitMQ      *configx.RabbitMQ      `mapstructure:"rabbitmq"`
	OpenTelemetry *configx.OpenTelemetry `mapstructure:"opentelemetry"`
	Postgres      *configx.Postgres      `mapstructure:"postgres"`
	Redis         *configx.Redis         `mapstructure:"redis"`
	CoinMarketCap *configx.CoinMarketCap `mapstructure:"coinmarketcap"`
}
