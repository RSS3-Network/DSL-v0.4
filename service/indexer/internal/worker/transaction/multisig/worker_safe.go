package multisig

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/safe"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/samber/lo"
)

// AddedOwner(address)
// 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26
func (m *MultiSign) handleGnosisSafeEventAddedOwner(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	event, err := m.gnosisSafeFilterer.ParseAddedOwner(log)
	if err != nil {
		return nil, fmt.Errorf("parse AddedOwner event: %w", err)
	}

	multiSig := metadata.MultiSig{
		Action: filter.ActionAddOwner,
		Owner:  lo.ToPtr(event.Owner.String()),
	}

	transfer, err := m.buildGnosisSafeTransfer(ctx, transaction, log, multiSig)
	if err != nil {
		return nil, fmt.Errorf("build transfer: %w", err)
	}

	return transfer, nil
}

// RemovedOwner(address)
// 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf
func (m *MultiSign) handleGnosisSafeEventRemovedOwner(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	event, err := m.gnosisSafeFilterer.ParseRemovedOwner(log)
	if err != nil {
		return nil, fmt.Errorf("parse RemovedOwner event: %w", err)
	}

	multiSig := metadata.MultiSig{
		Action: filter.ActionRemoveOwner,
		Owner:  lo.ToPtr(event.Owner.String()),
	}

	transfer, err := m.buildGnosisSafeTransfer(ctx, transaction, log, multiSig)
	if err != nil {
		return nil, fmt.Errorf("build transfer: %w", err)
	}

	return transfer, nil
}

// ApproveHash(bytes32,address)
// 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c

// ChangedFallbackHandler(address)
// 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0

// ChangedGuard(address)
// 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2

// ChangedThreshold(uint256)
// 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93
func (m *MultiSign) handleGnosisSafeEventChangedThreshold(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	event, err := m.gnosisSafeFilterer.ParseChangedThreshold(log)
	if err != nil {
		return nil, fmt.Errorf("parse ChangedThreshold event: %w", err)
	}

	multiSig := metadata.MultiSig{
		Action:    filter.ActionChangeThreshold,
		Threshold: event.Threshold,
	}

	transfer, err := m.buildGnosisSafeTransfer(ctx, transaction, log, multiSig)
	if err != nil {
		return nil, fmt.Errorf("build transfer: %w", err)
	}

	return transfer, nil
}

// DisabledModule(address)
// 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276

// EnabledModule(address)
// 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440

// ExecutionFromModuleFailure(address)
// 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375

// ExecutionFromModuleSuccess(address)
// 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8

// SafeReceived(address,uint256)
// 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d

// SignMsg(bytes32)
// 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4

type ExecutionTransactionInput struct {
	To             common.Address
	Value          *big.Int
	Data           []byte
	Operation      uint8
	SafeTxGas      *big.Int
	BaseGas        *big.Int
	GasPrice       *big.Int
	GasToken       common.Address
	RefundReceiver common.Address
	Signatures     []byte
}

// ExecutionSuccess(bytes32,uint256)
// 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e
func (m *MultiSign) handleGnosisSafeEventExecutionSuccess(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	_, err := m.gnosisSafeFilterer.ParseExecutionSuccess(log)
	if err != nil {
		return nil, fmt.Errorf("parse ExecutionSuccess event: %w", err)
	}

	isRejectionExecution, err := m.isRejectionExecution(transaction)
	if err != nil {
		return nil, fmt.Errorf("is rejection execution: %w", err)
	}

	multiSig := metadata.MultiSig{
		Action: func() string {
			if isRejectionExecution {
				return filter.ActionRejection
			}

			return filter.ActionExecution
		}(),
		Success: false,
	}

	transfer, err := m.buildGnosisSafeTransfer(ctx, transaction, log, multiSig)
	if err != nil {
		return nil, fmt.Errorf("build transfer: %w", err)
	}

	return transfer, nil
}

// ExecutionFailure(bytes32,uint256)
// 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23
func (m *MultiSign) handleGnosisSafeEventExecutionFailure(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	_, err := m.gnosisSafeFilterer.ParseExecutionFailure(log)
	if err != nil {
		return nil, fmt.Errorf("parse ExecutionFailure event: %w", err)
	}

	isRejectionExecution, err := m.isRejectionExecution(transaction)
	if err != nil {
		return nil, fmt.Errorf("is rejection execution: %w", err)
	}

	multiSig := metadata.MultiSig{
		Action: func() string {
			if isRejectionExecution {
				return filter.ActionRejection
			}

			return filter.ActionExecution
		}(),
		Success: false,
	}

	transfer, err := m.buildGnosisSafeTransfer(ctx, transaction, log, multiSig)
	if err != nil {
		return nil, fmt.Errorf("build transfer: %w", err)
	}

	return transfer, nil
}

func (m *MultiSign) isRejectionExecution(transaction model.Transaction) (bool, error) {
	var sourceData ethereum.SourceData
	if err := json.Unmarshal(transaction.SourceData, &sourceData); err != nil {
		return false, fmt.Errorf("unmarshal source data: %w", err)
	}

	method, exists := m.gnosisSafeABI.Methods["execTransaction"]
	if !exists {
		return false, fmt.Errorf("method not found: %s", "execTransaction")
	}

	// Get the input data of the transaction
	data := sourceData.Transaction.Data()

	if len(data) <= 4 {
		return false, fmt.Errorf("invalid data length: %d", len(data))
	}

	var input ExecutionTransactionInput
	values, err := method.Inputs.UnpackValues(data[4:])
	if err != nil {
		return false, fmt.Errorf("unpack input values: %w", err)
	}

	if err := method.Inputs.Copy(&input, values); err != nil {
		return false, fmt.Errorf("copy input values: %w", err)
	}

	// If the amount of the native token transfer is 0 and the data is empty,
	// it is considered to be a transaction used to reject transactions from other owners
	return input.Value.Cmp(big.NewInt(0)) == 0 && len(input.Data) == 0, nil
}

// SafeSetup(address,address[],uint256,address,address)
// 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8
func (m *MultiSign) handleGnosisSafeEventSafeSetup(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	_, err := m.gnosisSafeFilterer.ParseSafeSetup(log)
	if err != nil {
		return nil, fmt.Errorf("parse SafeSetup event: %w", err)
	}

	multiSig := metadata.MultiSig{
		Action: filter.ActionCreate,
	}

	transfer, err := m.buildGnosisSafeTransfer(ctx, transaction, log, multiSig)
	if err != nil {
		return nil, fmt.Errorf("build transfer: %w", err)
	}

	return transfer, nil
}

func (m *MultiSign) buildGnosisSafeTransfer(ctx context.Context, transaction model.Transaction, log types.Log, multiSig metadata.MultiSig) (*model.Transfer, error) {
	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("unsupported network %s: %w", transaction.Network, err)
	}

	caller, err := safe.NewGnosisSafeCaller(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new gnosis safe caller: %w", err)
	}

	callOptions := bind.CallOpts{
		// BlockNumber: new(big.Int).SetUint64(log.BlockNumber), // TODO Need archive node
		Context: ctx,
	}

	owners, err := caller.GetOwners(&callOptions)
	if err != nil {
		return nil, fmt.Errorf("get owners: %w", err)
	}

	threshold, err := caller.GetThreshold(&callOptions)
	if err != nil {
		return nil, fmt.Errorf("get threshold: %w", err)
	}

	version, err := caller.VERSION(&callOptions)
	if err != nil {
		return nil, fmt.Errorf("get version: %w", err)
	}

	vault := metadata.Vault{
		Address: log.Address.String(),
		Owners: lo.Map(owners, func(owner common.Address, i int) string {
			return owner.String()
		}),
		Threshold: threshold,
		Version:   version,
	}

	multiSig.Vault = vault

	metadataRaw, err := json.Marshal(multiSig)
	if err != nil {
		return nil, fmt.Errorf("marshal metadata: %w", err)
	}

	transfer := model.Transfer{
		TransactionHash: transaction.Hash,
		Timestamp:       transaction.Timestamp,
		BlockNumber:     big.NewInt(transaction.BlockNumber),
		Tag:             filter.TagTransaction,
		Type:            filter.TransactionMultiSig,
		Index:           int64(log.Index),
		AddressFrom:     transaction.AddressFrom,
		AddressTo:       transaction.AddressTo,
		Metadata:        metadataRaw,
		Network:         transaction.Network,
		Platform:        protocol.PlatformGnosisSafe,
		RelatedUrls:     ethereum.BuildURL([]string{}, ethereum.BuildScanURL(transaction.Network, transaction.Hash)),
	}

	return &transfer, nil
}
