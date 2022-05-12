package server

import (
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
)

var _ command.Interface = &Server{}

type Server struct {
	config         *config.Config
	databaseClient *database.Client
}

func (s *Server) Initialize() error {
	return nil
}

func (s *Server) Run() error {
	if err := s.Initialize(); err != nil {
		return err
	}

	return nil
}

func New(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}
