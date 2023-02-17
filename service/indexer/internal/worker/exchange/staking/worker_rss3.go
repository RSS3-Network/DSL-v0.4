package staking

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/database/model/transaction"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/rss3"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/shopspring/decimal"
)

func (s *Staking) handleRSS3StakingDeposited(ctx context.Context, _transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	if log.Address != rss3.AddressStaking {
		return nil, UnsupportedPlatform
	}

	filterer, err := rss3.NewStakingFilterer(log.Address, nil)
	if err != nil {
		return nil, fmt.Errorf("new rss3 staking filterer: %w", err)
	}

	event, err := filterer.ParseDeposited(log)
	if err != nil {
		return nil, fmt.Errorf("parse Deposited event: %w", err)
	}

	period := transaction.Period{
		Start: _transaction.Timestamp.String(),
		End:   _transaction.Timestamp.Add(time.Duration(event.Duration.Int64()) * time.Second).String(),
	}

	return s.buildRSS3StakingTransfer(ctx, _transaction, log, rss3.AddressStaking, event.Amount, &period, filter.ActionStakingDeposit)
}

func (s *Staking) handleRSS3StakingWithdrawn(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	if log.Address != rss3.AddressStaking {
		return nil, UnsupportedPlatform
	}

	filterer, err := rss3.NewStakingFilterer(log.Address, nil)
	if err != nil {
		return nil, fmt.Errorf("new rss3 staking filterer: %w", err)
	}

	event, err := filterer.ParseWithdrawn(log)
	if err != nil {
		return nil, fmt.Errorf("parse Withdrawn event: %w", err)
	}

	return s.buildRSS3StakingTransfer(ctx, transaction, log, rss3.AddressStaking, event.Amount, nil, filter.ActionStakingWithdraw)
}

func (s *Staking) handleRSS3StakingRewardsClaimed(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	if log.Address != rss3.AddressStaking {
		return nil, UnsupportedPlatform
	}

	filterer, err := rss3.NewStakingFilterer(log.Address, nil)
	if err != nil {
		return nil, fmt.Errorf("new rss3 staking filterer: %w", err)
	}

	event, err := filterer.ParseRewardsClaimed(log)
	if err != nil {
		return nil, fmt.Errorf("parse RewardsClaimed event: %w", err)
	}

	return s.buildRSS3StakingTransfer(ctx, transaction, log, rss3.AddressStaking, event.RewardAmount, nil, filter.ActionStakingCollect)
}

func (s *Staking) buildRSS3StakingTransfer(ctx context.Context, _transaction model.Transaction, log types.Log, tokenAddress common.Address, tokenValue *big.Int, period *transaction.Period, action string) (*model.Transfer, error) {
	var (
		token *metadata.Token
		err   error
	)

	if tokenAddress == ethereum.AddressGenesis {
		token, err = s.tokenClient.NatvieToMetadata(ctx, _transaction.Network)
	} else {
		token, err = s.tokenClient.ERC20ToMetadata(ctx, _transaction.Network, tokenAddress.String())
	}

	if err != nil {
		return nil, fmt.Errorf("get token metadata: %w", err)
	}

	token.SetValue(decimal.NewFromBigInt(tokenValue, 0))

	stakingMetadata := transaction.Staking{
		Action: action,
		Token:  *token,
		Period: period,
	}

	metadataRaw, err := json.Marshal(stakingMetadata)
	if err != nil {
		return nil, fmt.Errorf("marshal staking metadata: %w", err)
	}

	transfer := model.Transfer{
		TransactionHash: _transaction.Hash,
		Timestamp:       _transaction.Timestamp,
		BlockNumber:     big.NewInt(_transaction.BlockNumber),
		Tag:             filter.TagExchange,
		Type:            filter.ExchangeStaking,
		Index:           int64(log.Index),
		AddressFrom:     _transaction.AddressFrom,
		AddressTo:       _transaction.AddressTo,
		Metadata:        metadataRaw,
		Network:         _transaction.Network,
		Platform:        protocol.PlatformRSS3,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(_transaction.Network, _transaction.Hash)),
	}

	return &transfer, nil
}
