package auction

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	kurora "github.com/naturalselectionlabs/kurora/client"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/element"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/shopspring/decimal"

	"go.uber.org/zap"
)

var _ worker.Worker = (*internal)(nil)

const (
	SourceName = "auction"
	ThreadSize = 200
)

type internal struct {
	tokenClient  *token.Client
	kuroraClient *kurora.Client
}

func (i *internal) Name() string {
	return SourceName
}

func (i *internal) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
	}
}

func (i *internal) Initialize(ctx context.Context) error {
	return nil
}

func (i *internal) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) ([]model.Transaction, error) {
	var (
		resTransactions = make([]model.Transaction, 0)
		limiter         = make(chan struct{}, ThreadSize)

		waitGroup sync.WaitGroup
		locker    sync.Mutex
	)

	for _, platform := range platformList {

		limiter <- struct{}{}
		waitGroup.Add(1)

		go func(auctionPlatform string) {
			defer func() {
				<-limiter
				waitGroup.Done()
			}()

			var (
				internalTransactions []model.Transaction
				err                  error
			)

			switch auctionPlatform {
			case protocol.PlatformNouns:
				internalTransactions, err = i.handleNounsTransactions(ctx, message)
			case protocol.PlatformFoundation:
				internalTransactions, err = i.handleFoundationTransactions(ctx, message)
			case protocol.PlatformZora:
				internalTransactions, err = i.handleZoraTransactions(ctx, message)
			default:
				return
			}

			if err != nil {
				zap.L().Warn("handle transactions", zap.Error(err), zap.String("platform", auctionPlatform))

				return
			}

			locker.Lock()
			resTransactions = append(resTransactions, internalTransactions...)
			locker.Unlock()
		}(platform)
	}

	waitGroup.Wait()

	return resTransactions, nil
}

func (i *internal) buildCost(ctx context.Context, network string, address common.Address, value *big.Int) (*metadata.Token, error) {
	var costToken metadata.Token

	if address == ethereum.AddressGenesis || address == element.AddressNativeToken {
		nativeToken, err := i.tokenClient.Native(ctx, network)
		if err != nil {
			return nil, err
		}

		costValue := decimal.NewFromBigInt(value, 0)
		costValueDisplay := costValue.Shift(-int32(nativeToken.Decimals))

		costToken = metadata.Token{
			Name:         nativeToken.Name,
			Symbol:       nativeToken.Symbol,
			Decimals:     nativeToken.Decimals,
			Standard:     protocol.TokenStandardNative,
			Image:        nativeToken.Logo,
			Value:        &costValue,
			ValueDisplay: &costValueDisplay,
		}
	} else {
		erc20Token, err := i.tokenClient.ERC20(ctx, network, address.String())
		if err != nil {
			return nil, err
		}

		costValue := decimal.NewFromBigInt(value, 0)
		costValueDisplay := costValue.Shift(-int32(erc20Token.Decimals))

		costToken = metadata.Token{
			Name:            erc20Token.Name,
			Symbol:          erc20Token.Symbol,
			Decimals:        erc20Token.Decimals,
			Standard:        protocol.TokenStandardERC20,
			ContractAddress: erc20Token.ContractAddress,
			Image:           erc20Token.Logo,
			Value:           &costValue,
			ValueDisplay:    &costValueDisplay,
		}
	}

	return &costToken, nil
}

func (i *internal) Jobs() []worker.Job {
	return nil
}

func New(kuroraClient *kurora.Client) worker.Worker {
	return &internal{
		tokenClient:  token.New(),
		kuroraClient: kuroraClient,
	}
}
