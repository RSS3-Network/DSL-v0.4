package configx

import (
	"fmt"

	"github.com/naturalselectionlabs/pregod/common/protocol"
)

type Mode string

type File interface{}

type HTTP struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

var _ fmt.Stringer = &Postgres{}

type Postgres struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Database    string `mapstructure:"database"`
	SocialDBURL string `mapstructure:"socialdburl"`
}

func (p *Postgres) String() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		p.User, p.Password, p.Host, p.Port, p.Database,
	)
}

func (p *Postgres) SocialString() string {
	return p.SocialDBURL
}

var _ fmt.Stringer = &PostgresEtl{}

type PostgresEtl struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func (p *PostgresEtl) String() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		p.User, p.Password, p.Host, p.Port, p.Database,
	)
}

var _ fmt.Stringer = (*Kurora)(nil)

type Kurora struct {
	Endpoint string `mapstructure:"endpoint"`
}

func (k *Kurora) String() string {
	return k.Endpoint
}

var _ fmt.Stringer = &RabbitMQ{}

type RabbitMQ struct {
	TLS       bool   `mapstructure:"tls"`
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	QueueWork string `mapstructure:"queue_work"`
}

func (r *RabbitMQ) String() string {
	protocol := "amqp"
	if r.TLS {
		protocol = "amqps"
	}
	return fmt.Sprintf(
		"%s://%s:%s@%s:%d",
		protocol, r.User, r.Password, r.Host, r.Port,
	)
}

var _ fmt.Stringer = &OpenTelemetry{}

type OpenTelemetry struct {
	Enabled bool   `mapstructure:"enabled"`
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
	Path    string `mapstructure:"path"`
}

func (o *OpenTelemetry) String() string {
	return fmt.Sprintf(
		"http://%s:%d%s",
		o.Host, o.Port, o.Path,
	)
}

type Moralis struct {
	Key string `mapstructure:"key"`
}

type Redis struct {
	Addr       string `mapstructure:"addr"`
	Password   string `mapstructure:"password"`
	DB         int    `mapstructure:"db"`
	TLSEnabled bool   `mapstructure:"tls"`
}

type CoinMarketCap struct {
	APIKey string `mapstructure:"apikey"`
}

type Infura struct {
	ProjectID string `mapstructure:"projectid"`
}

type EIP1577 struct {
	Endpoint string `mapstructure:"endpoint"`
}

type NFTScan struct {
	APIKey string `mapstructure:"apiKey"`
}

type Gateway struct {
	EthEndpoint string `mapstructure:"ethendpoint"`
}

type IPFS struct {
	IO       string `mapstructure:"io"`
	Internal string `mapstructure:"internal"`
}

type RPC struct {
	General   RPCNetwork     `mapstructure:"general"`
	Alchemy   AlchemyNetwork `mapstructure:"alchemy"`
	PregodETL RPCNetwork     `mapstructure:"pregod_etl"`
	IPFS      IPFS           `mapstructure:"ipfs"`
}

type RPCNetwork struct {
	Ethereum          *RPCEndpoint `mapstructure:"ethereum"`
	Polygon           *RPCEndpoint `mapstructure:"polygon"`
	BinanceSmartChain *RPCEndpoint `mapstructure:"bsc"`
	XDAI              *RPCEndpoint `mapstructure:"xdai"`
	Crossbell         *RPCEndpoint `mapstructure:"crossbell"`
	Arbitrum          *RPCEndpoint `mapstructure:"arbitrum"`
	Optimism          *RPCEndpoint `mapstructure:"optimism"`
	Avalanche         *RPCEndpoint `mapstructure:"avalanche"`
	Celo              *RPCEndpoint `mapstructure:"celo"`
	Fantom            *RPCEndpoint `mapstructure:"fantom"`
}

type Sound struct {
	APIKey string `mapstructure:"apikey"`
}

func (r RPCNetwork) network2EP() map[string]*RPCEndpoint {
	return map[string]*RPCEndpoint{
		protocol.NetworkEthereum:          r.Ethereum,
		protocol.NetworkPolygon:           r.Polygon,
		protocol.NetworkBinanceSmartChain: r.BinanceSmartChain,
		protocol.NetworkXDAI:              r.XDAI,
		protocol.NetworkCrossbell:         r.Crossbell,
		protocol.NetworkArbitrum:          r.Arbitrum,
		protocol.NetworkOptimism:          r.Optimism,
		protocol.NetworkAvalanche:         r.Avalanche,
		protocol.NetworkCelo:              r.Celo,
		protocol.NetworkFantom:            r.Fantom,
	}
}

func (r RPCNetwork) HTTP(network string) string {
	mapNetwork := r.network2EP()
	if mapNetwork[network] == nil {
		return ""
	}
	return mapNetwork[network].HTTP
}

func (r RPCNetwork) WebSocket(network string) string {
	mapNetwork := r.network2EP()
	if mapNetwork[network] == nil {
		return ""
	}
	return mapNetwork[network].WebSocket
}

type RPCEndpoint struct {
	HTTP      string `mapstructure:"http"`
	WebSocket string `mapstructure:"websocket"`
}

type AlchemyNetwork struct {
	Ethereum string `mapstructure:"ethereum"`
	Polygon  string `mapstructure:"polygon"`
	Arbitrum string `mapstructure:"arbitrum"`
	Optimism string `mapstructure:"optimism"`
}
