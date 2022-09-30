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
	Mode          configx.Mode           `json:"mode"`
	HTTP          *configx.HTTP          `mapstructure:"http"`
	RabbitMQ      *configx.RabbitMQ      `mapstructure:"rabbitmq"`
	OpenTelemetry *configx.OpenTelemetry `mapstructure:"opentelemetry"`
	Postgres      *configx.Postgres      `mapstructure:"postgres"`
	Redis         *configx.Redis         `mapstructure:"redis"`
	CoinMarketCap *configx.CoinMarketCap `mapstructure:"coinmarketcap"`
	RPC           *configx.RPC           `mapstructure:"rpc"`
}

var ConfigHub Config

func Initialize() {
	viper.SetConfigName("hub")
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

	if err := viper.Unmarshal(&ConfigHub); err != nil {
		logrus.Fatalln(err)
	}

	var exporter trace.SpanExporter
	var err error

	if ConfigHub.OpenTelemetry == nil {
		if exporter, err = opentelemetry.DialWithPath(opentelemetry.DefaultPath); err != nil {
			loggerx.Global().Error("opentelemetry DialWithPath failed", zap.Error(err))
		}
	} else if ConfigHub.OpenTelemetry.Enabled {
		if exporter, err = opentelemetry.DialWithURL(ConfigHub.OpenTelemetry.String()); err != nil {
			loggerx.Global().Error("opentelemetry DialWithURL failed", zap.Error(err))
		}
	}

	otel.SetTracerProvider(trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("pregod-1-1-hub"),
			semconv.ServiceVersionKey.String(protocol.Version),
		)),
	))
}
