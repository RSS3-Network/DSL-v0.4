package server

import (
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	// "gorm.io/gorm"
)

var _ command.Interface = &Server{}

type Server struct {
	config *config.Config
	// databaseClient *gorm.DB
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
