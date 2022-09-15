package config

import (
	"strings"

	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Config struct {
	Mode          configx.Mode           `json:"mode"`
	HTTP          *configx.HTTP          `mapstructure:"http"`
	RabbitMQ      *configx.RabbitMQ      `mapstructure:"rabbitmq"`
	OpenTelemetry *configx.OpenTelemetry `mapstructure:"opentelemetry"`
	Postgres      *configx.Postgres      `mapstructure:"postgres"`
	Redis         *configx.Redis         `mapstructure:"redis"`
	CoinMarketCap *configx.CoinMarketCap `mapstructure:"coinmarketcap"`
	RPC           *configx.RPC           `mapstructure:"rpc"`

	DatabaseClient *gorm.DB
}

var ConfigBeat Config

func Initialize() {
	viper.SetConfigName("hub") // use the same config
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

	if err := viper.Unmarshal(&ConfigBeat); err != nil {
		logrus.Fatalln(err)
	}

	var err error
	ConfigBeat.DatabaseClient, err = database.Dial(ConfigBeat.Postgres.String(), false)
	if err != nil {
		panic(err)
	}
}
