package crossbell

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	contract "github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/crossbell"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/config"
	"github.com/naturalselectionlabs/pregod/service/crawler/internal/crawler"
	lop "github.com/samber/lo/parallel"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

var (
	_ crawler.Crawler = (*service)(nil)
)

type service struct {
}

func New(config *config.Config) crawler.Crawler {
	return &service{}
}

func (s *service) Name() string {
	return protocol.NetworkCrossbell
}

func (s *service) Run() error {
	ctx := context.Background()
	addressContract := []common.Address{contract.AddressCharacter, contract.AddressLinkList}

	lop.ForEach(addressContract, func(contract common.Address, i int) {
		transactions, err := s.GetTransactions(ctx, contract)
		if err != nil {
			loggerx.Global().Error("crossbell: GetTransactions error", zap.Error(err))

			return
		}
	})
}

func (s *service) GetTransactions(ctx context.Context, contract common.Address) ([]*model.Transaction, error) {
	tracer := otel.Tracer("crossbell")
	_, trace := tracer.Start(ctx, "crossbell:GetTransactions")
	var internalTransactions []*model.Transaction
	var err error
	defer func() { opentelemetry.Log(trace, nil, internalTransactions, err) }()

}
