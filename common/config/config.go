package configx

import "fmt"

type File interface{}

type HTTP struct {
	Host string
	Port int
}

var _ fmt.Stringer = &Postgres{}

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func (p *Postgres) String() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		p.User, p.Password, p.Host, p.Port, p.Database,
	)
}

var _ fmt.Stringer = &RabbitMQ{}

type RabbitMQ struct {
	Host     string
	Port     int
	User     string
	Password string
}

func (r *RabbitMQ) String() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%d",
		r.User, r.Password, r.Host, r.Port,
	)
}

var _ fmt.Stringer = &OpenTelemetry{}

type OpenTelemetry struct {
	Host string
	Port int
	Path string
}

func (o *OpenTelemetry) String() string {
	return fmt.Sprintf(
		"http://%s:%d%s",
		o.Host, o.Port, o.Path,
	)
}

type Moralis struct {
	Key string
}
