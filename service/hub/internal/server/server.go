package server

import (
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/ipfs"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/handler"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/middlewarex"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/service"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/validatorx"
	"go.uber.org/zap"
)

var _ command.Interface = &Server{}

type Server struct {
	httpServer  *echo.Echo
	httpHandler *handler.Handler
	logger      *zap.Logger
}

func (s *Server) Initialize() (err error) {
	s.logger, _ = zap.NewProduction()

	databaseClient, err := database.Dial(config.ConfigHub.Postgres.String(), false)
	if err != nil {
		panic(err)
	}
	database.ReplaceGlobal(databaseClient)

	redisClient, err := cache.Dial(config.ConfigHub.Redis)
	if err != nil {
		return err
	}

	cache.ReplaceGlobal(redisClient)

	ipfs.New(config.ConfigHub.RPC.IPFS.IO)

	ethereumClientMap, err := ethclientx.Dial(config.ConfigHub.RPC, []string{protocol.NetworkEthereum, protocol.NetworkPolygon, protocol.NetworkCrossbell, protocol.NetworkBinanceSmartChain})
	if err != nil {
		return err
	}

	for network, client := range ethereumClientMap {
		ethclientx.ReplaceGlobal(network, client)
	}

	svc := service.New()
	s.httpHandler = handler.New(svc)

	s.httpServer = echo.New()
	s.httpServer.HideBanner = true
	s.httpServer.HidePort = true
	s.httpServer.HTTPErrorHandler = handler.ErrorFunc
	s.httpServer.Validator = validatorx.Default

	s.httpServer.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	s.httpServer.Use(middlewarex.ZapLogger(s.logger))

	s.httpServer.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":        "ok",
			"version":       fmt.Sprintf("%s-%s", protocol.Version, protocol.Build),
			"documentation": "https://docs.rss3.io/PreGod/api/",
		})
	})

	// GET
	s.httpServer.GET("/notes/:address", s.httpHandler.GetNotesFunc, middlewarex.APIMiddleware)
	s.httpServer.GET("/assets/:address", s.httpHandler.GetAssetsFunc, middlewarex.APIMiddleware)
	s.httpServer.GET("/exchanges/:exchange_type", s.httpHandler.GetExchangeListFunc)
	s.httpServer.GET("/platforms/:platform_type", s.httpHandler.GetPlatformListFunc)
	s.httpServer.GET("/profiles/:address", s.httpHandler.GetProfilesFunc, middlewarex.APIMiddleware)
	s.httpServer.GET("/ns/:address", s.httpHandler.GetNameResolveFunc)

	// LensResolveFunc temporary function to resolve Lens for Pinata
	s.httpServer.GET("/nslens/:address", s.httpHandler.LensResolveFunc)

	// POST
	s.httpServer.POST("/notes", s.httpHandler.BatchGetNotesFunc, middlewarex.CheckAPIKeyMiddleware)
	s.httpServer.POST("/profiles", s.httpHandler.BatchGetProfilesFunc, middlewarex.CheckAPIKeyMiddleware)

	// API KEY
	s.httpServer.POST("/apikey/apply", s.httpHandler.PostAPIKeyFunc)
	s.httpServer.GET("/apikey", s.httpHandler.GetAPIKeyFunc)

	return nil
}

func (s *Server) Run() error {
	if err := s.Initialize(); err != nil {
		return err
	}

	return s.httpServer.Start(net.JoinHostPort(config.ConfigHub.HTTP.Host, strconv.Itoa(config.ConfigHub.HTTP.Port)))
}
