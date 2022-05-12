package config

type Config struct {
	RabbitMQ *RabbitMQ
	Moralis  *Moralis
}

type RabbitMQ struct {
	URL string
}

type Moralis struct {
	Key string
}
