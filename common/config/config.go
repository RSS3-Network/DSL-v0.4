package configx

import "fmt"

type File interface{}

type HTTP struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

var _ fmt.Stringer = &Postgres{}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func (p *Postgres) String() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		p.User, p.Password, p.Host, p.Port, p.Database,
	)
}

var _ fmt.Stringer = &RabbitMQ{}

type RabbitMQ struct {
	TLS      bool   `mapstructure:"tls"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

func (r *RabbitMQ) String() string {
	protocol := "amqp"
	if r.TLS {
		protocol = "amqps"
	}
	return fmt.Sprintf(
		"%s://%s:%s@%s:%d",
		protocol, r.User, r.Password, r.Host, r.Port,
	)
}

var _ fmt.Stringer = &OpenTelemetry{}

type OpenTelemetry struct {
	Enabled bool   `mapstructure:"enabled"`
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	Path    string `mapstructure:"path"`
}

func (o *OpenTelemetry) String() string {
	return fmt.Sprintf(
		"http://%s:%d%s",
		o.Host, o.Port, o.Path,
	)
}

type Moralis struct {
	Key string `mapstructure:"key"`
}

type Redis struct {
	Addr       string `mapstructure:"addr"`
	Password   string `mapstructure:"password"`
	DB         int    `mapstructure:"db"`
	TLSEnabled bool   `mapstructure:"tls"`
}

type CoinMarketCap struct {
	APIKey string `mapstructure:"apikey"`
}
