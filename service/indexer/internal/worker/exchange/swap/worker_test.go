package swap_test

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/naturalselectionlabs/pregod/common/database/model/metadata"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/uniswap"
	"github.com/naturalselectionlabs/pregod/common/datasource/ethereum/contract/zerox"
	"github.com/shopspring/decimal"
)

var (
	AddressETH = common.HexToAddress("0x0000000000000000000000000000000000001010")

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

func TestName(t *testing.T) {
	ethereumClient, err := ethclient.Dial("https://damp-delicate-glitter.quiknode.pro/22ae6184085ebd6115988ae67da9fb099645aa9f/")
	if err != nil {
		t.Fatalf("failed to connect to ethereum client: %v", err)
	}

	transaction, _, err := ethereumClient.TransactionByHash(context.Background(), common.HexToHash("0x99ce39ef831d9587118dc26caf3a2ae7e313a9ffb56779bf27780991aa7757cc"))
	if err != nil {
		t.Fatalf("failed to get transaction: %v", err)
	}

	if transaction.To() == nil {
		t.Errorf("not is a swap transaction")

		return
	}

	if _, exist := RouterMap[*transaction.To()]; !exist {
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

	swapMetadata := metadata.Swap{}

	for token, value := range tokenMap {
		t.Log(token, value)

		erc20Contract, err := erc20.NewERC20(token, ethereumClient)
		if err != nil {
			t.Fatalf("failed to get erc20 contract: %v", err)
		}

		tokenSymbol, err := erc20Contract.Symbol(&bind.CallOpts{})
		if err != nil {
			t.Fatalf("failed to get token symbol: %v", err)
		}

		tokenDecimals, err := erc20Contract.Decimals(&bind.CallOpts{})
		if err != nil {
			t.Fatalf("failed to get token decimals: %v", err)
		}

		tokenValue := decimal.NewFromBigInt(value, 0).DivRound(decimal.NewFromBigInt(big.NewInt(1), int32(tokenDecimals)), int32(tokenDecimals))

		switch value.Cmp(big.NewInt(0)) {
		case -1:
			swapMetadata.TokenFrom = metadata.SwapToken{
				Address:  token.String(),
				Symbol:   tokenSymbol,
				Decimals: tokenDecimals,
				Value:    tokenValue.Abs(),
			}
		case 0:
			continue
		case 1:
			swapMetadata.TokenTo = metadata.SwapToken{
				Address:  token.String(),
				Symbol:   tokenSymbol,
				Decimals: tokenDecimals,
				Value:    tokenValue,
			}
		}
	}

	data, err := json.MarshalIndent(swapMetadata, "", "\t")
	if err != nil {
		t.Fatalf("failed to marshal swap metadata: %v", err)
	}

	t.Log(string(data))
}
