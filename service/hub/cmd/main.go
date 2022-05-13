package main

import (
	"net"

	"github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/constant"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCommand = cobra.Command{
	Use:           "hub",
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       constant.Version,
}

func main() {
	srv := server.New(&config.Config{
		HTTP: &configx.HTTP{
			Host: net.IPv4zero.String(),
			Port: 80,
		},
		RabbitMQ: &configx.RabbitMQ{
			Host:     "localhost",
			Port:     5672,
			User:     "guest",
			Password: "guest",
		},
		OpenTelemetry: &configx.OpenTelemetry{
			Host: "localhost",
			Port: 14268,
			Path: "/api/traces",
		},
		Postgres: &configx.Postgres{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "password",
			Database: "postgres",
		},
	})

	rootCommand.RunE = func(cmd *cobra.Command, args []string) error {
		return srv.Run()
	}

	if err := rootCommand.Execute(); err != nil {
		logrus.Fatalln(err)
	}
}
