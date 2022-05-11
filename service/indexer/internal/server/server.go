package server

import (
	"github.com/naturalselectionlabs/pregod/common/command"
	"github.com/naturalselectionlabs/pregod/common/database"
)

var _ command.Interface = &Server{}

type Server struct {
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
