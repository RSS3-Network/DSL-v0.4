package ens_test

import (
	"strings"
	"testing"
	"time"

	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/worker/ens"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
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

func loadConfig() {
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
}

var client *ens.Client

var address = "0x827431510a5D249cE4fdB7F00C83a3353F471848"
var ensDomain = "henryqw.eth"
var expiry = time.Unix(1956691715, 0)

func init() {
	loadConfig()
	client = ens.New(ConfigHub.RPC.General.Ethereum.HTTP)
}
func TestGetProfile(t *testing.T) {
	result, err := client.GetProfile(address)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetReverseResolve(t *testing.T) {
	result, err := client.GetReverseResolve(address)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, result, ensDomain)
}

func TestGetResolvedAddress(t *testing.T) {
	result, err := client.GetResolvedAddress(ensDomain)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, result, address)
}

func TestGetENSExpiry(t *testing.T) {
	result, err := client.GetENSExpiry(ensDomain)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, result, expiry)
}
