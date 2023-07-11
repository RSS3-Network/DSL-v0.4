package server

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/metadata_url"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/handler"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/handler/doc"
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

	ethDbClient, err := database.Dial(config.ConfigHub.EthereumEtl.String(), false)
	if err != nil {
		return err
	}

	database.ReplaceEthDb(ethDbClient)

	databaseClient, err := database.Dial(config.ConfigHub.Postgres.String(), false)
	if err != nil {
		panic(err)
	}
	database.ReplaceGlobal(databaseClient)

	// Compatible with old configuration
	if config.ConfigHub.Postgres.SocialDBURL != "" {
		socialDatabaseClient, err := database.Dial(config.ConfigHub.Postgres.SocialString(), false)
		if err != nil {
			panic(err)
		}
		database.ReplaceSocial(socialDatabaseClient)
	} else {
		database.ReplaceSocial(databaseClient)
	}

	redisClient, err := cache.Dial(config.ConfigHub.Redis)
	if err != nil {
		return err
	}

	cache.ReplaceGlobal(redisClient)

	metadata_url.New(config.ConfigHub.RPC.IPFS.IO)

	ethereumClientMap, err := ethclientx.Dial(config.ConfigHub.RPC, protocol.EthclientNetworks)
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
		u, _ := url.JoinPath(c.Request().URL.Path, "openapi")
		return c.JSON(http.StatusOK, map[string]string{
			"status":        "ok",
			"version":       fmt.Sprintf("%s-%s", protocol.Version, protocol.Build),
			"documentation": u,
		})
	})

	// OpenAPI doc for endpoints
	s.httpServer.GET("/openapi", doc.New().OpenAPI())

	// GET
	s.httpServer.GET(handler.PathGetNotes, s.httpHandler.GetNotesFunc, middlewarex.APIMiddleware)
	s.httpServer.GET(handler.PathGetAssets, s.httpHandler.GetAssetsFunc, middlewarex.APIMiddleware)
	s.httpServer.GET(handler.PathGetExchanges, s.httpHandler.GetExchangeListFunc)
	s.httpServer.GET(handler.PathGetPlatformList, s.httpHandler.GetPlatformListFunc)
	s.httpServer.GET(handler.PathGetProfiles, s.httpHandler.GetProfilesFunc2)
	s.httpServer.GET(handler.PathGetNameResolve, s.httpHandler.GetNameResolveFunc)
	s.httpServer.GET(handler.PathGetTransaction, s.httpHandler.GetTransactionByHashFunc)

	// ActivityPub Mastodon
	s.httpServer.GET(handler.PathGetMastodon, s.httpHandler.GetMastodonFunc)

	// POST
	s.httpServer.POST(handler.PathBatchGetSocialNotes, s.httpHandler.BatchGetSocialNotesFunc, middlewarex.CheckAPIKeyMiddleware)
	s.httpServer.POST(handler.PathBatchGetNotes, s.httpHandler.BatchGetNotesFunc, middlewarex.CheckAPIKeyMiddleware)
	s.httpServer.POST(handler.PathBatchGetProfiles, s.httpHandler.BatchGetProfilesFunc2, middlewarex.CheckAPIKeyMiddleware)

	// End of year Wrapped
	s.httpServer.GET("/wrapped/:address", s.httpHandler.GetWrappedFunc)

	// API KEY
	s.httpServer.POST(handler.PathPostAPIKey, s.httpHandler.PostAPIKeyFunc)
	s.httpServer.GET(handler.PathGetAPIKey, s.httpHandler.GetAPIKeyFunc)

	// WS Initialize
	// go svc.WsHub.Run()
	// s.httpServer.GET("/ws/notes", s.httpHandler.GetNotesWsFunc)

	return nil
}

func (s *Server) Run() error {
	if err := s.Initialize(); err != nil {
		return err
	}

	return s.httpServer.Start(net.JoinHostPort(config.ConfigHub.HTTP.Host, strconv.Itoa(config.ConfigHub.HTTP.Port)))
}
