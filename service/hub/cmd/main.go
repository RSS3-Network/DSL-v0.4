package main

import (
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/logger"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var rootCommand = cobra.Command{
	Use:           "hub",
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       protocol.Version,
}

func main() {
	config.Initialize()

	if err := logger.Initialize(string(config.ConfigHub.Mode)); err != nil {
		logrus.Fatalln(err)
	}

	srv := server.New(&config.ConfigHub)

	rootCommand.RunE = func(cmd *cobra.Command, args []string) error {
		return srv.Run()
	}

	if err := rootCommand.Execute(); err != nil {
		logger.Global().Fatal("hub execution failed", zap.Error(err))
	}
}
