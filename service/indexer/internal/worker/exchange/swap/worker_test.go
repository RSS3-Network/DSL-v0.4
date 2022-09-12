package swap_test

import (
	"context"
	"encoding/json"
	"github.com/naturalselectionlabs/pregod/common/ethclientx"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	configx "github.com/naturalselectionlabs/pregod/common/config"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/uniswap"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/zerox"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/internal/token"
	"github.com/shopspring/decimal"
)

var (
	RouterMap = map[common.Address]string{
		common.HexToAddress("0x1111111254fb6c44bAC0beD2854e76F90643097d"): "1inch v4",
		common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"): "Uniswap V3",
		common.HexToAddress("0x4cb18386e5d1f34dc6eea834bf3534a970a3f8e7"): "MEV Bot",
		common.HexToAddress("0x98c3d3183c4b8a650614ad179a1a98be0a8d6b8e"): "MEV Bot",
		common.HexToAddress("0x10ed43c718714eb63d5aa57b78b54704e256024e"): "PancakeSwap: Router v2",
		common.HexToAddress("0x1b02da8cb0d097eb8d57a175b88c7d8b47997506"): "SushiSwap: Router",
		common.HexToAddress("0x881d40237659c251811cec9c364ef91dc08d300c"): "Metamask: Swap Router",
		common.HexToAddress("0xdef1c0ded9bec7f1a1670819833240f027b25eff"): "0x: Exchange Proxy",
	}
)

func TestBuildTransferMetadata(t *testing.T) {
	ethereumClientMap, err := ethclientx.Dial(&configx.RPC{
		General: configx.RPCNetwork{
			Ethereum: &configx.RPCEndpoint{
				HTTP: "https://rpc.rss3.dev/networks/ethereum",
			},
		},
	}, []string{protocol.NetworkEthereum})
	if err != nil {
		t.Fatalf("failed to create ethereum client: %v", err)
	}

	ethereumClient, exists := ethereumClientMap[protocol.NetworkEthereum]
	if !exists {
		t.Fatalf("failed to find ethereum client")
	}

	transaction, _, err := ethereumClient.TransactionByHash(context.Background(), common.HexToHash("0x4b37b3dcb3014a9f09d0ce9ced7385a9662f0464b738c5957a96cd8c4af12160"))
	if err != nil {
		t.Fatalf("failed to get transaction: %v", err)
	}

	if transaction.To() == nil {
		t.Errorf("not is a swap transaction")

		return
	}

	router, exist := RouterMap[*transaction.To()]
	if !exist {
		t.Errorf("not is a swap transaction")

		return
	}

	receipt, err := ethereumClient.TransactionReceipt(context.Background(), transaction.Hash())
	if err != nil {
		t.Fatalf("failed to get transaction receipt: %v", err)
	}

	tokenMap := map[common.Address]*big.Int{}

	if *transaction.To() == common.HexToAddress("0xdef1c0ded9bec7f1a1670819833240f027b25eff") {
		for _, log := range receipt.Logs {
			switch log.Topics[0] {
			case zerox.EventHashTransformedERC20:
				zeroXContact, err := zerox.NewZeroX(common.HexToAddress(log.Address.Hex()), ethereumClient)
				if err != nil {
					t.Fatalf("falied to create zerox contract: %v", err)
				}

				event, err := zeroXContact.ParseTransformedERC20(*log)
				if err != nil {
					t.Fatalf("falied to parse transformed erc20 event: %v", err)
				}

				token0Value, exist := tokenMap[event.InputToken]
				if !exist {
					token0Value = big.NewInt(0)
				}

				token1Value, exist := tokenMap[event.OutputToken]
				if !exist {
					token1Value = big.NewInt(0)
				}

				tokenMap[event.InputToken] = token0Value.Sub(token0Value, event.InputTokenAmount)
				tokenMap[event.OutputToken] = token1Value.Add(token1Value, event.OutputTokenAmount)
			default:
				continue
			}
		}
	} else {
		for _, log := range receipt.Logs {
			switch log.Topics[0] {
			case uniswap.EventHashSwapV2:
				uniswapPoolContact, err := uniswap.NewPoolV2(common.HexToAddress(log.Address.Hex()), ethereumClient)
				if err != nil {
					t.Fatalf("failed to create uniswap pool contract: %v", err)
				}

				event, err := uniswapPoolContact.ParseSwap(*log)
				if err != nil {
					t.Fatalf("failed to parse swap event: %v", err)
				}

				token0, err := uniswapPoolContact.Token0(&bind.CallOpts{})
				if err != nil {
					t.Fatalf("failed to get token0: %v", err)
				}

				token1, err := uniswapPoolContact.Token1(&bind.CallOpts{})
				if err != nil {
					t.Fatalf("failed to get token1: %v", err)
				}

				token0Value, exist := tokenMap[token0]
				if !exist {
					token0Value = big.NewInt(0)
				}

				token1Value, exist := tokenMap[token1]
				if !exist {
					token1Value = big.NewInt(0)
				}

				if event.Amount0In.Cmp(big.NewInt(0)) == 1 {
					// Swap token0 to token1
					tokenMap[token0] = token0Value.Sub(token0Value, event.Amount0In)
					tokenMap[token1] = token1Value.Add(token1Value, event.Amount1Out)
				} else {
					// Swap token1 to token0
					tokenMap[token0] = token0Value.Add(token0Value, event.Amount0Out)
					tokenMap[token1] = token1Value.Sub(token1Value, event.Amount1In)
				}
			case uniswap.EventHashSwapV3:
				uniswapPoolContact, err := uniswap.NewPoolV3(common.HexToAddress(log.Address.Hex()), ethereumClient)
				if err != nil {
					t.Fatalf("failed to create uniswap pool contract: %v", err)
				}

				event, err := uniswapPoolContact.ParseSwap(*log)
				if err != nil {
					t.Fatalf("failed to parse swap event: %v", err)
				}

				token0, err := uniswapPoolContact.Token0(&bind.CallOpts{})
				if err != nil {
					t.Fatalf("failed to get token0: %v", err)
				}

				token1, err := uniswapPoolContact.Token1(&bind.CallOpts{})
				if err != nil {
					t.Fatalf("failed to get token1: %v", err)
				}

				token0Value, exist := tokenMap[token0]
				if !exist {
					token0Value = big.NewInt(0)
				}

				token1Value, exist := tokenMap[token1]
				if !exist {
					token1Value = big.NewInt(0)
				}

				t.Logf("token0: %v, token1: %v, amount0: %v, amount1: %v", token0, token1, event.Amount0, event.Amount1)

				tokenMap[token0] = token0Value.Sub(token0Value, event.Amount0)
				tokenMap[token1] = token1Value.Sub(token1Value, event.Amount1)
			default:
				continue
			}
		}
	}

	swapMetadata := metadata.Swap{
		Protocol: router,
	}

	for internalToken, value := range tokenMap {
		t.Log(internalToken, value)

		tokenClient := token.New(ethereumClientMap)
		erc20, err := tokenClient.ERC20(context.Background(), protocol.NetworkEthereum, strings.ToLower(internalToken.String()))
		if err != nil {
			t.Fatalf("failed to get erc20 contract: %v", err)
		}

		tokenValueTo := decimal.NewFromBigInt(value, 0)
		tokenValueDisplayTo := tokenValueTo.Shift(-int32(erc20.Decimals))

		tokenValueFrom := tokenValueTo.Abs()
		tokenValueDisplayFrom := tokenValueFrom.Shift(-int32(erc20.Decimals))

		switch value.Cmp(big.NewInt(0)) {
		case -1:
			swapMetadata.TokenFrom = metadata.Token{
				Name:            erc20.Name,
				Image:           erc20.Logo,
				Symbol:          erc20.Symbol,
				Decimals:        erc20.Decimals,
				Standard:        protocol.TokenStandardERC20,
				ContractAddress: erc20.ContractAddress,
				Value:           &tokenValueFrom,
				ValueDisplay:    &tokenValueDisplayFrom,
			}
		case 0:
			continue
		case 1:
			swapMetadata.TokenTo = metadata.Token{
				Name:            erc20.Name,
				Image:           erc20.Logo,
				Symbol:          erc20.Symbol,
				Decimals:        erc20.Decimals,
				Standard:        protocol.TokenStandardERC20,
				ContractAddress: erc20.ContractAddress,
				Value:           &tokenValueTo,
				ValueDisplay:    &tokenValueDisplayTo,
			}
		}
	}

	data, err := json.MarshalIndent(swapMetadata, "", "\t")
	if err != nil {
		t.Fatalf("failed to marshal swap metadata: %v", err)
	}

	t.Log(string(data))
}
