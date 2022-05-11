package main

import (
	"net"

	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCommand = cobra.Command{
	Use:           "hub",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func main() {
	srv := server.New(&config.Config{
		HTTP: &config.HTTP{
			Host: net.IPv4zero.String(),
			Port: 80,
		},
	})

	rootCommand.RunE = func(cmd *cobra.Command, args []string) error {
		return srv.Run()
	}

	if err := rootCommand.Execute(); err != nil {
		logrus.Fatalln(err)
	}
}
