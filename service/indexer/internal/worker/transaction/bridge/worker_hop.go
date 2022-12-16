package bridge

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/hop"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
	"github.com/samber/lo"
)

var hopTokenMap = map[common.Address]*common.Address{
	hop.AddressBridgeL1EthereumUSDC:     lo.ToPtr(common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")),
	hop.AddressBridgeL1EthereumUSDT:     lo.ToPtr(common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")),
	hop.AddressBridgeL1EthereumMATIC:    lo.ToPtr(common.HexToAddress("0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0")),
	hop.AddressBridgeL1EthereumDAI:      lo.ToPtr(common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f")),
	hop.AddressBridgeL1EthereumETH:      nil,
	hop.AddressBridgeL1EthereumHOP:      lo.ToPtr(common.HexToAddress("0xc5102fe9359fd9a28f877a67e36b0f050d81a3cc")),
	hop.AddressBridgeL1EthereumSNX:      lo.ToPtr(common.HexToAddress("0xc011a73ee8576fb46f5e1c5751ca3b9fe0af2a6f")),
	hop.AddressBridgeL2GnosisHOP:        lo.ToPtr(common.HexToAddress("0xc5102fe9359fd9a28f877a67e36b0f050d81a3cc")),
	hop.AddressBridgeL2PolygonHOP:       lo.ToPtr(common.HexToAddress("0xc5102fE9359FD9a28f877a67E36B0F050d81a3CC")),
	hop.AddressBridgeL2OptimismHOP:      lo.ToPtr(common.HexToAddress("0xc5102fe9359fd9a28f877a67e36b0f050d81a3cc")),
	hop.AddressBridgeL2ArbitrumHOP:      lo.ToPtr(common.HexToAddress("0xc5102fE9359FD9a28f877a67E36B0F050d81a3CC")),
	hop.AddressAMMWrapperL2GnosisUSDC:   lo.ToPtr(common.HexToAddress("0xddafbb505ad214d7b80b1f830fccc89b60fb7a83")),
	hop.AddressAMMWrapperL2GnosisUSDT:   lo.ToPtr(common.HexToAddress("0x4ecaba5870353805a9f068101a40e0f32ed605c6")),
	hop.AddressAMMWrapperL2GnosisMATIC:  lo.ToPtr(common.HexToAddress("0x7122d7661c4564b7c6cd4878b06766489a6028a2")),
	hop.AddressAMMWrapperL2GnosisDAI:    lo.ToPtr(common.HexToAddress("0xe91d153e0b41518a2ce8dd3d7944fa863463a97d")),
	hop.AddressAMMWrapperL2GnosisETH:    lo.ToPtr(common.HexToAddress("0x6a023ccd1ff6f2045c3309768ead9e68f978f6e1")),
	hop.AddressAMMWrapperL2PolygonUSDC:  lo.ToPtr(common.HexToAddress("0x2791bca1f2de4661ed88a30c99a7a9449aa84174")),
	hop.AddressAMMWrapperL2PolygonUSDT:  lo.ToPtr(common.HexToAddress("0xc2132d05d31c914a87c6611c10748aeb04b58e8f")),
	hop.AddressAMMWrapperL2PolygonMATIC: lo.ToPtr(common.HexToAddress("0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270")),
	hop.AddressAMMWrapperL2PolygonDAI:   lo.ToPtr(common.HexToAddress("0x8f3cf7ad23cd3cadbd9735aff958023239c6a063")),
	hop.AddressAMMWrapperL2PolygonETH:   lo.ToPtr(common.HexToAddress("0x7ceb23fd6bc0add59e62ac25578270cff1b9f619")),
	hop.AddressAMMWrapperL2OptimismUSDC: lo.ToPtr(common.HexToAddress("0x7f5c764cbc14f9669b88837ca1490cca17c31607")),
	hop.AddressAMMWrapperL2OptimismUSDT: lo.ToPtr(common.HexToAddress("0x94b008aa00579c1307b0ef2c499ad98a8ce58e58")),
	hop.AddressAMMWrapperL2OptimismDAI:  lo.ToPtr(common.HexToAddress("0xda10009cbd5d07dd0cecc66161fc93d7c9000da1")),
	hop.AddressAMMWrapperL2OptimismETH:  lo.ToPtr(common.HexToAddress("0x4200000000000000000000000000000000000006")),
	hop.AddressAMMWrapperL2ArbitrumUSDC: lo.ToPtr(common.HexToAddress("0xff970a61a04b1ca14834a43f5de4533ebddb5cc8")),
	hop.AddressAMMWrapperL2ArbitrumUSDT: lo.ToPtr(common.HexToAddress("0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9")),
	hop.AddressAMMWrapperL2ArbitrumDAI:  lo.ToPtr(common.HexToAddress("0xda10009cbd5d07dd0cecc66161fc93d7c9000da1")),
	hop.AddressAMMWrapperL2ArbitrumETH:  nil,
}

func (w *Worker) handleHopTransferSentToL2(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	contract, err := hop.NewL1Bridge(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("get l1 bridge address: %w", err)
	}

	event, err := contract.ParseTransferSentToL2(log)
	if err != nil {
		return nil, fmt.Errorf("parse transfer sent event: %w", err)
	}

	tokenAddress, exists := hopTokenMap[common.HexToAddress(transaction.AddressTo)]
	if !exists {
		return nil, fmt.Errorf("token address not found for network: %s", transaction.Network)
	}

	return w.buildTransfer(ctx, transaction, log, common.HexToAddress(transaction.AddressFrom), event.Recipient, platform, event.ChainId.Uint64(), tokenAddress, event.Amount, filter.TransactionBridge)
}

func (w *Worker) handleHopTransferSent(ctx context.Context, transaction model.Transaction, log types.Log) (*model.Transfer, error) {
	platform, exists := platformMap[log.Address]
	if !exists {
		return nil, UnsupportedPlatform
	}

	ethereumClient, err := ethclientx.Global(transaction.Network)
	if err != nil {
		return nil, fmt.Errorf("ethclientx global: %w", err)
	}

	contract, err := hop.NewL2AMMWapper(log.Address, ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new l2 bridge filterer: %w", err)
	}

	event, err := contract.ParseTransferSent(log)
	if err != nil {
		return nil, fmt.Errorf("parse transfer sent event: %w", err)
	}

	tokenAddress, exists := hopTokenMap[common.HexToAddress(transaction.AddressTo)]
	if !exists {
		return nil, fmt.Errorf("token address not found for network: %s", transaction.Network)
	}

	return w.buildTransfer(ctx, transaction, log, common.HexToAddress(transaction.AddressFrom), event.Recipient, platform, event.ChainId.Uint64(), tokenAddress, event.Amount, filter.TransactionBridge)
}
