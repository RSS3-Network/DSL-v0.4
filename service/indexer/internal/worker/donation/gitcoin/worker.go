package gitcoin

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/pregod/common/database"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/donation"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/gitcoin"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/gitcoin/quadratic"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/gitcoin/round"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/metadata_url"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/common/utils/loggerx"
	"github.com/naturalselectionlabs/pregod/common/utils/opentelemetry"
	"github.com/naturalselectionlabs/pregod/common/worker/zksync"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/donation/gitcoin/job"
	"github.com/shopspring/decimal"

	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

const (
	Name                          = protocol.PlatformGitcoin
	ContractAddressPolygon        = "0xb99080b9407436ebb2b8fe56d45ffa47e9bb8877"
	ContractAddressEth            = "0x7d655c57f71464b6f83811c55d84009cd9f5221c"
	ContractAddressEthereumNative = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
)

var _ worker.Worker = (*service)(nil)

type service struct {
	tokenClient  *token.Client
	zksyncClient *zksync.Client
}

func (s *service) Name() string {
	return Name
}

func (s *service) Networks() []string {
	return []string{
		protocol.NetworkEthereum,
		protocol.NetworkPolygon,
		protocol.NetworkZkSync,
	}
}

func (s *service) Initialize(ctx context.Context) error {
	return nil
}

func (s *service) Handle(ctx context.Context, message *protocol.Message, transactions []model.Transaction) (internalTransactions []model.Transaction, err error) {
	tracer := otel.Tracer("gitcoin_worker")
	_, trace := tracer.Start(ctx, "gitcoin_worker:Handle")

	defer func() { opentelemetry.Log(trace, transactions, internalTransactions, err) }()

	internalTransactionMap := make(map[string]model.Transaction)

	for _, transaction := range transactions {
		var transfers []model.Transfer
		var err error

		switch {
		case message.Network == protocol.NetworkEthereum && transaction.AddressTo == ContractAddressEth:
			transfers, err = s.handleGitcoinEthereum(ctx, message, transaction)
		case message.Network == protocol.NetworkEthereum && common.HexToAddress(transaction.AddressTo) == gitcoin.AddressRound:
			transfers, err = s.handleGitcoinEthereumBetaRound(ctx, message, transaction)
		case message.Network == protocol.NetworkPolygon && transaction.AddressTo == ContractAddressPolygon:
			transfers, err = s.handleGitcoinEthereum(ctx, message, transaction)
		case message.Network == protocol.NetworkZkSync:
			transfers, err = s.handlerGitcoinZkSync(ctx, message, transaction)
		default:
			internalTransactionMap[transaction.Hash] = transaction

			continue
		}

		if err != nil {
			loggerx.Global().Error("failed to handle gitcoin transaction", zap.Error(err))

			continue
		}

		for _, transfer := range transfers {
			// Copy the transaction to map
			value, exist := internalTransactionMap[transaction.Hash]
			if !exist {
				value = transaction

				// Ignore transfers data that will not be updated
				value.Transfers = make([]model.Transfer, 0)
			}

			value.Tag, value.Type = filter.UpdateTagAndType(transfer.Tag, value.Tag, transfer.Type, value.Type)
			value.Transfers = append(value.Transfers, transfer)

			if value.Tag == filter.TagDonation {
				value.Platform = Name
			}

			value.Owner = transaction.AddressFrom

			internalTransactionMap[transaction.Hash] = value
		}
	}

	internalTransactions = make([]model.Transaction, 0)

	for _, internalTransaction := range internalTransactionMap {
		internalTransactions = append(internalTransactions, internalTransaction)
	}

	return internalTransactions, nil
}

func (s *service) handleGitcoinEthereum(ctx context.Context, message *protocol.Message, transaction model.Transaction) (transfers []model.Transfer, err error) {
	tracer := otel.Tracer("worker_gitcoin")
	_, span := tracer.Start(ctx, "worker_gitcoin:handleGitcoin")

	defer opentelemetry.Log(span, transaction, transfers, err)

	// Unsupported network
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal source data: %w", err)
	}

	for _, log := range sourceData.Receipt.Logs {
		// Filter anonymous log
		if len(log.Topics) == 0 {
			continue
		}

		if log.Topics[0] != gitcoin.EventHashDonationSent {
			continue
		}

		gitcoinContract, err := gitcoin.NewGitcoin(log.Address, ethclient)
		if err != nil {
			loggerx.Global().Error("get gitcoin contract error", zap.Error(err))

			continue
		}

		event, err := gitcoinContract.ParseDonationSent(*log)
		if err != nil {
			loggerx.Global().Error("parse donation sent event error", zap.Error(err))

			continue
		}

		if event.Dest == ethereum.AddressGenesis {
			continue
		}

		var project donation.GitcoinProject

		if err := database.Global().
			Model((*donation.GitcoinProject)(nil)).
			Where(map[string]any{
				"admin_address": strings.ToLower(event.Dest.Hex()),
			}).
			Order("id DESC").
			First(&project).
			Error; err != nil {
			loggerx.Global().Error("get gitcoin project error", zap.Error(err))

			continue
		}

		sourceData, err := json.Marshal(log)
		if err != nil {
			loggerx.Global().Error("marshal source data error", zap.Error(err))

			continue
		}

		transfer := model.Transfer{
			TransactionHash: transaction.Hash,
			Timestamp:       transaction.Timestamp,
			BlockNumber:     big.NewInt(transaction.BlockNumber),
			Tag:             transaction.Tag,
			Index:           int64(log.Index),
			AddressFrom:     strings.ToLower(event.Donor.String()),
			AddressTo:       strings.ToLower(event.Dest.String()),
			Metadata:        metadata.Default,
			Network:         transaction.Network,
			Source:          transaction.Source,
			SourceData:      sourceData,
		}

		transfer.RelatedUrls = append(transfer.RelatedUrls, fmt.Sprintf("https://gitcoin.co/grants/%d/%s", project.ID, project.Slug))

		var tokenMetadata *metadata.Token

		if strings.ToLower(event.Token.String()) == ContractAddressEthereumNative {
			tokenMetadata, err = s.tokenClient.NatvieToMetadata(ctx, transaction.Network)
		} else {
			tokenMetadata, err = s.tokenClient.ERC20ToMetadata(ctx, transaction.Network, event.Token.String())
		}
		if err != nil {
			loggerx.Global().Error("get token error", zap.Error(err))
		}

		tokenValue := decimal.NewFromBigInt(event.Amount, 0)
		tokenValueDisplay := tokenValue.Shift(-int32(tokenMetadata.Decimals))
		tokenMetadata.Value = &tokenValue
		tokenMetadata.ValueDisplay = &tokenValueDisplay

		metadataRaw, err := json.Marshal(&metadata.Donation{
			ID:          project.ID,
			Title:       project.Title,
			Description: project.Description,
			Logo:        project.Logo,
			Platform:    protocol.PlatformGitcoin,
			Token:       *tokenMetadata,
		})
		if err != nil {
			loggerx.Global().Error("marshal metadata error", zap.Error(err))

			continue
		}

		transfer.Metadata = metadataRaw
		transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagDonation, transfer.Tag, filter.DonationDonate, transfer.Type)

		if transfer.Tag == filter.TagDonation {
			transfer.Platform = Name
		}

		transfers = append(transfers, transfer)
	}

	return transfers, nil
}

func (s *service) handlerGitcoinZkSync(ctx context.Context, message *protocol.Message, transaction model.Transaction) (transfers []model.Transfer, err error) {
	for _, transfer := range transaction.Transfers {
		if transfer.AddressTo == "" ||
			transfer.AddressTo == "0x0" ||
			transfer.AddressTo == "0x0000000000000000000000000000000000000000" {
			continue
		}

		var data zksync.GetTransactionData

		if err := json.Unmarshal(transfer.SourceData, &data); err != nil {
			continue
		}

		// gitcoin grant info
		var project donation.GitcoinProject

		if err := database.Global().
			Model((*donation.GitcoinProject)(nil)).
			Where(map[string]any{
				"admin_address": strings.ToLower(transfer.AddressTo),
			}).
			Order("id DESC").
			First(&project).
			Error; err != nil {
			loggerx.Global().Error("get gitcoin project error", zap.Error(err))

			continue
		}

		transfer.RelatedUrls = append(transfer.RelatedUrls, fmt.Sprintf("https://gitcoin.co/grants/%d/%s", project.ID, project.Slug))

		// token matedata
		tokenInfo, _, err := s.zksyncClient.GetToken(ctx, uint(data.Transaction.Operation.Token))
		if err != nil {
			loggerx.Global().Error("handlerGitcoinZkSync: failed to get token", zap.Error(err), zap.String("token_id", strconv.Itoa(int(data.Transaction.Operation.Token))))

			continue
		}

		var tokenMetadata *metadata.Token

		if tokenInfo.Address == ethereum.AddressGenesis.String() {
			tokenMetadata, err = s.tokenClient.NatvieToMetadata(ctx, message.Network)
		} else {
			tokenMetadata, err = s.tokenClient.ERC20ToMetadata(ctx, message.Network, tokenInfo.Address)
		}
		if err != nil {
			loggerx.Global().Error("get token error", zap.Error(err))

			continue
		}

		tokenValue, err := decimal.NewFromString(data.Transaction.Operation.Amount)
		if err != nil {
			loggerx.Global().Error("decimal NewFromString error", zap.Error(err))

			continue
		}

		tokenValueDisplay := tokenValue.Shift(-int32(tokenMetadata.Decimals))
		tokenMetadata.Value = &tokenValue
		tokenMetadata.ValueDisplay = &tokenValueDisplay

		// metadata
		metadataRaw, err := json.Marshal(&metadata.Donation{
			ID:          project.ID,
			Title:       project.Title,
			Description: project.Description,
			Logo:        project.Logo,
			Platform:    protocol.PlatformGitcoin,
			Token:       *tokenMetadata,
		})
		if err != nil {
			loggerx.Global().Error("marshal metadata error", zap.Error(err))

			continue
		}

		transfer.Metadata = metadataRaw
		transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagDonation, transfer.Tag, filter.DonationDonate, transfer.Type)

		if transfer.Tag == filter.TagDonation {
			transfer.Platform = Name
		}

		transfers = append(transfers, transfer)
	}

	return transfers, nil
}

func (s *service) handleGitcoinEthereumBetaRound(ctx context.Context, message *protocol.Message, transaction model.Transaction) (transfers []model.Transfer, err error) {
	tracer := otel.Tracer("worker_gitcoin")
	_, span := tracer.Start(ctx, "worker_gitcoin:handleGitcoin")

	defer opentelemetry.Log(span, transaction, transfers, err)

	// Unsupported network
	ethclient, err := ethclientx.Global(message.Network)
	if err != nil {
		return nil, err
	}

	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal source data: %w", err)
	}

	for _, log := range sourceData.Receipt.Logs {
		// Filter anonymous log
		if len(log.Topics) == 0 {
			continue
		}

		if log.Topics[0] != gitcoin.EventVoted {
			continue
		}

		quadraticContract, err := quadratic.NewQuadratic(log.Address, ethclient)
		if err != nil {
			loggerx.Global().Error("get Quadratic contract error", zap.Error(err))

			continue
		}

		voted, err := quadraticContract.ParseVoted(*log)
		if err != nil {
			loggerx.Global().Error("parse voted event error", zap.Error(err))

			continue
		}

		roundContract, err := round.NewRound(gitcoin.AddressRound, ethclient)
		if err != nil {
			loggerx.Global().Error("get Round contract error", zap.Error(err))

			continue
		}

		application, err := roundContract.Applications(&bind.CallOpts{}, voted.ApplicationIndex)
		if err != nil {
			loggerx.Global().Error("get application metadata error", zap.Error(err))

			continue
		}

		appIpfsUrl := metadata_url.GetDirectURL(fmt.Sprintf("%s%s", "ipfs://", application.MetaPtr.Pointer))

		projectMeta, err := s.getEthereumBetaRoundMetadata(ctx, appIpfsUrl)
		if err != nil {
			loggerx.Global().Error("get project metadata error", zap.Error(err))

			continue
		}

		sourceData, err := json.Marshal(log)
		if err != nil {
			loggerx.Global().Error("marshal source data error", zap.Error(err))

			continue
		}

		transfer := model.Transfer{
			TransactionHash: transaction.Hash,
			Timestamp:       transaction.Timestamp,
			BlockNumber:     big.NewInt(transaction.BlockNumber),
			Tag:             transaction.Tag,
			Index:           int64(log.Index),
			AddressFrom:     strings.ToLower(voted.Voter.String()),
			AddressTo:       strings.ToLower(voted.GrantAddress.String()),
			Metadata:        metadata.Default,
			Network:         transaction.Network,
			Source:          transaction.Source,
			SourceData:      sourceData,
		}

		transfer.RelatedUrls = append(transfer.RelatedUrls, fmt.Sprintf("%s/%s/%s-%v", "https://explorer.gitcoin.co/#/round/1", strings.ToLower(voted.RoundAddress.String()), strings.ToLower(voted.RoundAddress.String()), voted.ApplicationIndex))

		var tokenMetadata *metadata.Token

		if strings.ToLower(voted.Token.String()) == ContractAddressEthereumNative || voted.Token == ethereum.AddressGenesis {
			tokenMetadata, err = s.tokenClient.NatvieToMetadata(ctx, transaction.Network)
		} else {
			tokenMetadata, err = s.tokenClient.ERC20ToMetadata(ctx, transaction.Network, voted.Token.String())
		}
		if err != nil {
			loggerx.Global().Error("get token error", zap.Error(err))
		}

		tokenValue := decimal.NewFromBigInt(voted.Amount, 0)
		tokenValueDisplay := tokenValue.Shift(-int32(tokenMetadata.Decimals))
		tokenMetadata.Value = &tokenValue
		tokenMetadata.ValueDisplay = &tokenValueDisplay

		metadataRaw, err := json.Marshal(&metadata.Donation{
			Title:       projectMeta.Application.Project.Title,
			Description: projectMeta.Application.Project.Description,
			Logo:        metadata_url.GetDirectURL(fmt.Sprintf("%s%s", "ipfs://", projectMeta.Application.Project.LogoImg)),
			Platform:    protocol.PlatformGitcoin,
			Token:       *tokenMetadata,
		})
		if err != nil {
			loggerx.Global().Error("marshal metadata error", zap.Error(err))

			continue
		}

		transfer.Metadata = metadataRaw
		transfer.Tag, transfer.Type = filter.UpdateTagAndType(filter.TagDonation, transfer.Tag, filter.DonationDonate, transfer.Type)

		if transfer.Tag == filter.TagDonation {
			transfer.Platform = Name
		}

		transfers = append(transfers, transfer)
	}

	return transfers, nil
}

func (s *service) getEthereumBetaRoundMetadata(ctx context.Context, url string) (*ApplicationResp, error) {
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("http client do request: %w", err)
	}

	var response *ApplicationResp
	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return nil, err
	}

	if len(response.Signature) == 0 {
		return nil, fmt.Errorf("response return an error: %w", nil)
	}

	return response, nil
}

func (s *service) Jobs() []worker.Job {
	return []worker.Job{
		&job.GitcoinProjectJob{},
		&job.GitcoinAllGrantJob{},
	}
}

func New() worker.Worker {
	return &service{
		tokenClient:  token.New(),
		zksyncClient: zksync.New(),
	}
}
