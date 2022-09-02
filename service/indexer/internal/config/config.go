package config

import (
	"strings"

	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/zap"
)

type Config struct {
	Mode          configx.Mode           `mapstructure:"mode"`
	RabbitMQ      *configx.RabbitMQ      `mapstructure:"rabbitmq"`
	Postgres      *configx.Postgres      `mapstructure:"postgres"`
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

	var exporter trace.SpanExporter
	var err error

	if ConfigIndexer.OpenTelemetry == nil {
		if exporter, err = opentelemetry.DialWithPath(opentelemetry.DefaultPath); err != nil {
			loggerx.Global().Fatal(" opentelemetry.DialWithPath failed", zap.Error(err))
			return
		}
	} else if ConfigIndexer.OpenTelemetry.Enabled {
		if exporter, err = opentelemetry.DialWithURL(ConfigIndexer.OpenTelemetry.String()); err != nil {
			loggerx.Global().Fatal(" opentelemetry.DialWithPath failed", zap.Error(err))
			return
		}
	}

	otel.SetTracerProvider(trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("pregod-1-1-indexer"),
			semconv.ServiceVersionKey.String(protocol.Version),
		)),
	))
}
