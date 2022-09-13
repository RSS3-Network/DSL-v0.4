package server

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/ipfs"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/utils/shedlock"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler/ens"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler/lens"
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

var _ command.Interface = &Server{}

type Server struct {
	config             *config.Config
	rabbitmqConnection *rabbitmq.Connection
	rabbitmqChannel    *rabbitmq.Channel
	rabbitmqQueue      rabbitmq.Queue
	crawlers           []crawler.Crawler
	employer           *shedlock.Employer
}

func (s *Server) Initialize() (err error) {
	var exporter trace.SpanExporter

	if s.config.OpenTelemetry == nil {
		if exporter, err = opentelemetry.DialWithPath(opentelemetry.DefaultPath); err != nil {
			return err
		}
	} else if s.config.OpenTelemetry.Enabled {
		if exporter, err = opentelemetry.DialWithURL(s.config.OpenTelemetry.String()); err != nil {
			return err
		}
	}

	otel.SetTracerProvider(trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("pregod-1-1-crawler"),
			semconv.ServiceVersionKey.String(protocol.Version),
		)),
	))

	databaseClient, err := database.Dial(s.config.Postgres.String(), true)
	if err != nil {
		return err
	}

	database.ReplaceGlobal(databaseClient)

	redisClient, err := cache.Dial(s.config.Redis)
	if err != nil {
		return err
	}

	cache.ReplaceGlobal(redisClient)

	ethereumClientMap, err := ethclientx.Dial(s.config.RPC, []string{protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkCrossbell})
	if err != nil {
		return err
	}

	for network, client := range ethereumClientMap {
		ethclientx.ReplaceGlobal(network, client)
	}

	ipfs.New(s.config.RPC.IPFS.Internal)

	s.rabbitmqConnection, err = rabbitmq.Dial(s.config.RabbitMQ.String())
	if err != nil {
		return err
	}

	s.rabbitmqChannel, err = s.rabbitmqConnection.Channel()
	if err != nil {
		return err
	}

	if err := s.rabbitmqChannel.ExchangeDeclare(
		protocol.ExchangeJob, "direct", true, false, false, false, nil,
	); err != nil {
		return err
	}

	if s.rabbitmqQueue, err = s.rabbitmqChannel.QueueDeclare(
		protocol.IndexerWorkQueue, false, false, false, false, nil,
	); err != nil {
		return err
	}

	if err := s.rabbitmqChannel.QueueBind(
		s.rabbitmqQueue.Name, protocol.IndexerWorkRoutingKey, protocol.ExchangeJob, false, nil,
	); err != nil {
		return err
	}

	s.employer = shedlock.New()

	s.crawlers = []crawler.Crawler{
		ens.New(s.rabbitmqChannel, s.employer, s.config),
		lens.New(s.config),
	}

	return nil
}

func (s *Server) Run() error {
	if err := s.Initialize(); err != nil {
		return err
	}

	s.employer.Start()

	for _, internalCrawler := range s.crawlers {
		go func(internalCrawler crawler.Crawler) {
			if err := internalCrawler.Run(); err != nil {
				logrus.Errorf("[crawler] run error: %v", err)
			}
		}(internalCrawler)
	}

	stopchan := make(chan os.Signal, 1)
	signal.Notify(stopchan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-stopchan
	s.employer.Stop()

	return nil
}

func New(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}
