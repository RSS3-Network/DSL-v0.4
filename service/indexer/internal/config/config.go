package config

type Config struct {
	RabbitMQ *RabbitMQ
}

type RabbitMQ struct {
	URL string
}
