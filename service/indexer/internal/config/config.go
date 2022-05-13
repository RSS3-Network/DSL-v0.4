package config

import "github.com/naturalselectionlabs/pregod/common/config"

type Config struct {
	RabbitMQ *configx.RabbitMQ
	Moralis  *configx.Moralis
}
