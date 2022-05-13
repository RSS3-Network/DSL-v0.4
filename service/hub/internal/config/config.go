package config

type Config struct {
	HTTP     *HTTP
	RabbitMQ *RabbitMQ
	Jaeger   *Jaeger
	Postgres *Postgres
}

type HTTP struct {
	Host string
	Port int
}

type RabbitMQ struct {
	URL string
}

type Jaeger struct {
	URL string
}

type Postgres struct {
	DSN string
}
