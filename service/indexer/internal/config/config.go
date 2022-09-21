package config

import (
	"strings"

	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Mode          configx.Mode           `mapstructure:"mode"`
	RabbitMQ      *configx.RabbitMQ      `mapstructure:"rabbitmq"`
	Postgres      *configx.Postgres      `mapstructure:"postgres"`
	EthereumEtl   *configx.PostgresEtl   `mapstructure:"ethereumetl"`
	OpenTelemetry *configx.OpenTelemetry `mapstructure:"opentelemetry"`
	Moralis       *configx.Moralis       `mapstructure:"moralis"`
	Redis         *configx.Redis         `mapstructure:"redis"`
	CoinMarketCap *configx.CoinMarketCap `mapstructure:"coinmarketcap"`
	Infura        *configx.Infura        `mapstructure:"infura"`
	RPC           *configx.RPC           `mapstructure:"rpc"`
}

var ConfigIndexer Config

func Initialize() {
	viper.SetConfigName("indexer")
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

	if err := viper.Unmarshal(&ConfigIndexer); err != nil {
		logrus.Fatalln(err)
	}
}
