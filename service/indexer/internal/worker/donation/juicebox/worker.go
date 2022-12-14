package juicebox

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/juicebox"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/ipfs"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"golang.org/x/sync/errgroup"
)

var _ worker.Worker = (*internal)(nil)

type internal struct {
	tokenClient *token.Client
}

func (i *internal) Name() string {
	return "juicebox"
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
	errorGroup, ctx := errgroup.WithContext(ctx)

	contractAddresses := []common.Address{
		juicebox.AddressETHPaymentTerminal,
		juicebox.AddressController,
		juicebox.AddressTiered721DelegateProjectDeployer,
	}

	transactions = lo.Filter(transactions, func(transaction model.Transaction, _ int) bool {
		return lo.ContainsBy(contractAddresses, func(address common.Address) bool {
			return strings.EqualFold(transaction.AddressTo, address.String())
		})
	})

	for index := range transactions {
		internalIndex := index

		errorGroup.Go(func() error {
			transfers, err := i.handle(ctx, transactions[internalIndex])
			if err != nil {
				return fmt.Errorf("handle: %w", err)
			}

			transaction := transactions[internalIndex]
			transaction.Transfers = transfers

			for _, transfer := range transaction.Transfers {
				transaction.Tag, transaction.Type = filter.UpdateTagAndType(filter.TagDonation, transaction.Tag, transfer.Type, transaction.Type)
			}

			transaction.Platform = protocol.PlatformJuiceBox

			transactions[internalIndex] = transaction

			return nil
		})
	}

	if err := errorGroup.Wait(); err != nil {
		return nil, fmt.Errorf("handle message: %w", err)
	}

	return transactions, nil
}

func (i *internal) handle(ctx context.Context, transaction model.Transaction) (transfers []model.Transfer, err error) {
	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return nil, fmt.Errorf("unmarshal source data: %w", err)
	}

	for _, log := range sourceData.Receipt.Logs {
		var internalTransfers []model.Transfer

		topic := log.Topics[0]

		switch {
		case topic == juicebox.EventPay && log.Address == juicebox.AddressETHPaymentTerminal:
			internalTransfers, err = i.handleEventPay(ctx, transaction, *log)
		case topic == juicebox.EventLaunchProject && log.Address == juicebox.AddressController:
			internalTransfers, err = i.handleEventLaunchProject(ctx, transaction, *log)
		default:
			continue
		}

		if err != nil {
			return nil, fmt.Errorf("handle event pay: %w", err)
		}

		transfers = append(transfers, internalTransfers...)
	}

	return transfers, nil
}

func (i *internal) handleEventPay(ctx context.Context, transaction model.Transaction, log types.Log) ([]model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("get ethereum client: %w", err)
	}

	filterer, err := juicebox.NewETHPaymentTerminalFilterer(juicebox.AddressETHPaymentTerminal, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("create filterer: %w", err)
	}

	event, err := filterer.ParsePay(log)
	if err != nil {
		return nil, fmt.Errorf("parse pay event: %w", err)
	}

	nativeToken, err := i.tokenClient.NatvieToMetadata(ctx, transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("get native token: %w", err)
	}

	nativeToken.SetValue(decimal.NewFromBigInt(event.Amount, 0))

	donationMetadata, err := i.buildMetadata(ctx, event.ProjectId, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("build metadata: %w", err)
	}

	donationMetadata.Token = nativeToken

	transfer, err := i.buildTransfer(ctx, transaction, log.Index, lo.FromPtr(donationMetadata), filter.DonationDonate)
	if err != nil {
		return nil, fmt.Errorf("build transfer: %w", err)
	}

	transfers := []model.Transfer{
		lo.FromPtr(transfer),
	}

	return transfers, nil
}

func (i *internal) handleEventLaunchProject(ctx context.Context, transaction model.Transaction, log types.Log) ([]model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("get ethereum client: %w", err)
	}

	filterer, err := juicebox.NewControllerFilterer(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("create filterer: %w", err)
	}

	event, err := filterer.ParseLaunchProject(log)
	if err != nil {
		return nil, fmt.Errorf("parse launch project event: %w", err)
	}

	donationMetadata, err := i.buildMetadata(ctx, event.ProjectId, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("build metadata: %w", err)
	}

	transfer, err := i.buildTransfer(ctx, transaction, log.Index, lo.FromPtr(donationMetadata), filter.DonationLaunch)
	if err != nil {
		return nil, fmt.Errorf("build transfer: %w", err)
	}

	transfers := []model.Transfer{
		lo.FromPtr(transfer),
	}

	return transfers, nil
}

func (i *internal) buildMetadata(ctx context.Context, projectID *big.Int, ethereumClient *ethclient.Client) (*metadata.Donation, error) {
	caller, err := juicebox.NewProjectsCaller(juicebox.AddressProjects, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("create caller: %w", err)
	}

	callOptions := &bind.CallOpts{}

	cid, err := caller.MetadataContentOf(callOptions, projectID, decimal.Zero.BigInt())
	if err != nil {
		return nil, fmt.Errorf("get prject %d metadata cid: %w", projectID, err)
	}

	ipfsURL := ipfs.BuildURL(cid)

	data, err := ipfs.GetFileByURL(ipfsURL)
	if err != nil {
		return nil, fmt.Errorf("get project metadata by url %s: %w", ipfsURL, err)
	}

	var project Metadata
	if err := json.Unmarshal(data, &project); err != nil {
		return nil, fmt.Errorf("unmarshal project: %w", err)
	}

	donation := metadata.Donation{
		ID:          int(projectID.Int64()),
		Title:       project.Name,
		Description: project.Description,
		Logo:        ipfs.GetDirectURL(project.LogoURI),
		Platform:    protocol.PlatformJuiceBox,
	}

	return &donation, nil
}

func (i *internal) buildTransfer(ctx context.Context, transaction model.Transaction, index uint, donationMetadata metadata.Donation, transferType string) (*model.Transfer, error) {
	metadataRaw, err := json.Marshal(donationMetadata)
	if err != nil {
		return nil, fmt.Errorf("marshal metadata: %w", err)
	}

	transfer := model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     new(big.Int).SetInt64(transaction.BlockNumber),
		Tag:             filter.TagDonation,
		Type:            transferType,
		Index:           int64(index),
		AddressFrom:     transaction.AddressFrom,
		AddressTo:       strings.ToLower(juicebox.AddressETHPaymentTerminal.String()),
		Metadata:        metadataRaw,
		Network:         transaction.Network,
		Platform:        protocol.PlatformJuiceBox,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}

	return &transfer, nil
}

func (i *internal) Jobs() []worker.Job {
	return nil
}

func New() worker.Worker {
	return &internal{
		tokenClient: token.New(),
	}
}
