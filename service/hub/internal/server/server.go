package server

import (
	"net"
	"strconv"

	"github.com/labstack/echo"
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/config"
)

var _ command.Interface = &Server{}

type Server struct {
	config         *config.Config
	httpServer     *echo.Echo
	databaseClient *database.Client
}

func (s *Server) Initialize() error {
	s.httpServer = echo.New()
	s.httpServer.HideBanner = true
	s.httpServer.HidePort = true

	return nil
}

func (s *Server) Run() error {
	if err := s.Initialize(); err != nil {
		return err
	}

	return s.httpServer.Start(net.JoinHostPort(s.config.HTTP.Host, strconv.Itoa(s.config.HTTP.Port)))
}

func New(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}
