package main

import (
	"strings"

	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCommand = cobra.Command{
	Use:           "crawler",
	SilenceUsage:  true,
	SilenceErrors: true,
	Version:       protocol.Version,
}

func init() {
	rootCommand.PersistentFlags().String("socialdb", "", "social db name")
	rootCommand.PersistentFlags().Int("socialredisdb", -1, "social redis db name")
}

func main() {
	viper.SetConfigName("crawler")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/pregod/")
	viper.AddConfigPath("$HOME/.pregod/")
	viper.AddConfigPath("./deploy/config/")
	// `opentelemetry.host` -> `CONFIG_ENV_OPENTELEMETRY_HOST`
	viper.SetEnvPrefix("CONFIG_ENV")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalln(err)
	}

	// read config
	configCrawler := config.Config{}

	if err := viper.Unmarshal(&configCrawler); err != nil {
		logrus.Fatalln(err)
	}

	srv := server.New(&configCrawler)

	rootCommand.RunE = func(cmd *cobra.Command, args []string) error {
		// config: db overridden by args
		socialDB, _ := rootCommand.Flags().GetString("socialdb")
		if socialDB != "" {
			configCrawler.Postgres.Database = socialDB
		}
		socialRedisDB, _ := rootCommand.Flags().GetInt("socialredisdb")
		if socialRedisDB != -1 {
			configCrawler.Redis.DB = socialRedisDB
		}

		return srv.Run()
	}

	if err := rootCommand.Execute(); err != nil {
		logrus.Fatalln(err)
	}
}
