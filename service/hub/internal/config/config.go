package config

type Config struct {
	HTTP     *HTTP
	RabbitMQ *RabbitMQ
}

type HTTP struct {
	Host string
	Port int
}

type RabbitMQ struct {
	URL string
}
