package crossbell

import (
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	"golang.org/x/net/context"
)

var (
	_ crawler.Crawler = (*service)(nil)
)

type service struct {
}

func New() crawler.Crawler {
	return &service{}
}

func (s *service) Name() string {
	return protocol.NetworkCrossbell
}

func (s *service) Run() error {
	ctx := context.Background()

	
}