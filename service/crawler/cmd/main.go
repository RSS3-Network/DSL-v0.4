package main

import (
	"github.com/naturalselectionlabs/pregod/common/constant"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCommand = cobra.Command{
	Use:           "crawler",
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       constant.Version,
}

func main() {
	srv := server.New(&config.Config{})

	rootCommand.RunE = func(cmd *cobra.Command, args []string) error {
		return srv.Run()
	}

	if err := rootCommand.Execute(); err != nil {
		logrus.Fatalln(err)
	}
}
