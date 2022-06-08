package coinmarketcap

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
	"github.com/naturalselectionlabs/pregod/common/cache"
	"github.com/naturalselectionlabs/pregod/service/indexer/internal/worker/token/contract"
	"github.com/shopspring/decimal"
)

const (
	Prefix = "https://pro-api.coinmarketcap.com"
)

var (
	client      = resty.New()
	Network2RPC = map[string]string{
		"ethereum":            "https://eth.rss3.dev",
		"binance_smart_chain": "https://bsc-dataseed.binance.org/",
		"polygon":             "https://polygon-mainnet.infura.io/v3/56ff3a0f0ad14ec19dc933038c69bfbe",
		"zksync":              "https://zksync2-testnet.zksync.dev",
		// "xdai":                "https://rpc.ankr.com/gnosis",
		"xdai":      "https://rpc.gnosischain.com/",
		"arbitrum":  "https://rpc.ankr.com/arbitrum",
		"avalanche": "https://rpc.ankr.com/avalanche",
		"celo":      "https://rpc.ankr.com/celo",
		"fantom":    "https://rpc.ankr.com/fantom",
		"harmony":   "https://rpc.ankr.com/harmony",
		"iotex":     "https://rpc.ankr.com/iotex",
		"moonbeam":  "https://rpc.ankr.com/moonbeam",
		"near":      "https://rpc.ankr.com/near",
		"nervos":    "https://rpc.ankr.com/nervos_gw",
		"syscoin":   "https://rpc.ankr.com/syscoin",
	}
	Network2Token = map[string]struct {
		Symbol  string
		Address string
	}{
		"ethereum":            {"ETH", ""},
		"binance_smart_chain": {"BNB", "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"},   // BNB address (on BSC)
		"polygon":             {"MATIC", "0x0000000000000000000000000000000000001010"}, // MATIC address (on polygon)
	}
)

func Init(apikey string) {
	client.SetHeaders(map[string]string{
		"Content-Type":      "application/json",
		"X-CMC_PRO_API_KEY": apikey,
	})
}

type CoinInfo struct {
	Logo                          string
	Name                          string
	Slug                          string
	Symbol                        string
	Category                      string
	Description                   string
	Decimals                      uint8    // TODO: cannot get from coinmarketcap
	SelfReportedCirculatingSupply *float64 `json:"self_reported_circulating_supply"`

	Supply *decimal.Decimal
}

type CoinInfos struct {
	Data map[string]*CoinInfo
}

func (i *CoinInfos) List() []*CoinInfo {
	result := []*CoinInfo{}
	for _, v := range i.Data {
		result = append(result, v)
	}
	return result
}

// Not the same as the documentation
type CoinInfosBySymbol struct {
	Data map[string][]*CoinInfo
}

func (i *CoinInfosBySymbol) List() []*CoinInfo {
	result := []*CoinInfo{}
	for _, v := range i.Data {
		result = append(result, v...)
	}
	return result
}

// https://coinmarketcap.com/api/documentation/v1/#operation/getV2CryptocurrencyInfo
// Will get a empty result when resp code is not OK
func getCoinInfo(ctx context.Context, network, address string) (*CoinInfo, error) {
	path := "/v2/cryptocurrency/info"

	result := CoinInfos{}
	resp, err := client.R().SetResult(&result).SetQueryParams(map[string]string{
		"address": address,
	}).Get(Prefix + path)
	if err != nil {
		return nil, err
	}
	if resp.IsError() { // skip this one and return a empty result
		info := new(CoinInfo)
		info.Supply = new(decimal.Decimal)
		return info, nil
		// return nil, fmt.Errorf("bad resp code: %v", resp.StatusCode())
	}
	info := result.List()[0]

	// get decimals
	decimals, err := getDecimals(ctx, network, address)
	if err != nil {
		return nil, err
	}
	info.Decimals = decimals

	// set supply default value
	var supply decimal.Decimal
	if info.SelfReportedCirculatingSupply != nil {
		supply = decimal.NewFromFloat(*info.SelfReportedCirculatingSupply)
	}
	info.Supply = &supply

	return info, nil
}

func getDecimals(ctx context.Context, network, addressHex string) (uint8, error) {
	rpcUrl := Network2RPC[network]
	if rpcUrl == "" {
		rpcUrl = Network2RPC["ethereum"]
	}

	client, err := ethclient.DialContext(ctx, rpcUrl)
	if err != nil {
		return 0, err
	}

	contractCaller, err := contract.NewERC20(common.HexToAddress(addressHex), client)
	if err != nil {
		return 0, err
	}

	return contractCaller.Decimals(&bind.CallOpts{})
}

func CachedGetCoinInfo(ctx context.Context, network, address string) (*CoinInfo, error) {
	key := fmt.Sprintf("getcoininfo.address.%s", address)
	ttl := time.Hour

	result := &CoinInfo{}
	exists, err := cache.GetMsgPack(ctx, key, result)
	if err != nil {
		return nil, err
	}

	if !exists {
		result, err = getCoinInfo(ctx, network, address)
		if err != nil {
			return nil, err
		}
		if err = cache.SetMsgPack(ctx, key, result, ttl); err != nil {
			return nil, err
		}
	}

	return result, nil
}

// for token native transfer
func getCoinInfoByNetwork(ctx context.Context, network string) (*CoinInfo, error) {
	token, exists := Network2Token[network]
	if !exists {
		return nil, errors.New("symbol for the network not exists")
	}

	path := "/v2/cryptocurrency/info"

	result := CoinInfosBySymbol{}
	resp, err := client.R().SetResult(&result).SetQueryParams(map[string]string{
		"symbol": token.Symbol,
	}).Get(Prefix + path)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("bad resp code: %v", resp.StatusCode())
	}
	info := result.List()[0]

	// get decimals
	switch {
	case token.Symbol == "ETH":
		info.Decimals = 18
	default:
		decimals, err := getDecimals(ctx, network, token.Address)
		if err != nil {
			return nil, err
		}
		info.Decimals = decimals
	}

	// set supply default value
	var supply decimal.Decimal
	if info.SelfReportedCirculatingSupply != nil {
		supply = decimal.NewFromFloat(*info.SelfReportedCirculatingSupply)
	}
	info.Supply = &supply

	return info, nil
}

// for token native transfer
func CachedGetCoinInfoByNetwork(ctx context.Context, network string) (*CoinInfo, error) {
	key := fmt.Sprintf("getcoininfobynetwork.network.%s", network)
	ttl := time.Hour

	result := &CoinInfo{}
	exists, err := cache.GetMsgPack(ctx, key, result)
	if err != nil {
		return nil, err
	}

	if !exists {
		result, err = getCoinInfoByNetwork(ctx, network)
		if err != nil {
			return nil, err
		}
		if err = cache.SetMsgPack(ctx, key, result, ttl); err != nil {
			return nil, err
		}
	}

	return result, nil
}
