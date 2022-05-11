package config

type Config struct {
	HTTP *HTTP
}

type HTTP struct {
	Host string
	Port int
}
