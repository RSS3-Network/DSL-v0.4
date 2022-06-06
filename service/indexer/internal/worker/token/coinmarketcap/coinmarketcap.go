package coinmarketcap

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
	"github.com/naturalselectionlabs/pregod/common/cache"
)

const (
	Prefix   = "https://pro-api.coinmarketcap.com"
	TokenABI = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"spender","type":"address"},{"name":"tokens","type":"uint256"}],"name":"approve","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"from","type":"address"},{"name":"to","type":"address"},{"name":"tokens","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"tokenOwner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"to","type":"address"},{"name":"tokens","type":"uint256"}],"name":"transfer","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"tokenOwner","type":"address"},{"name":"spender","type":"address"}],"name":"allowance","outputs":[{"name":"remaining","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"tokens","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"tokenOwner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"tokens","type":"uint256"}],"name":"Approval","type":"event"}]`
)

var (
	client      = resty.New()
	NetWork2RPC = map[string]string{
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

// https://coinmarketcap.com/api/documentation/v1/#operation/getV2CryptocurrencyInfo
func getCoinInfo(ctx context.Context, network, address string) (*CoinInfo, error) {
	path := "/v2/cryptocurrency/info"

	result := CoinInfos{}
	resp, err := client.R().SetResult(&result).SetQueryParams(map[string]string{
		"address": address,
	}).Get(Prefix + path)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("bad resp code: %v", resp.StatusCode())
	}
	info := result.List()[0]

	// get decimals
	decimals, err := getDecimals(ctx, network, address)
	if err != nil {
		return nil, err
	}
	info.Decimals = decimals

	return info, nil
}

func getDecimals(ctx context.Context, network, addressHex string) (uint8, error) {
	rpcUrl := NetWork2RPC[network]
	if rpcUrl == "" {
		rpcUrl = NetWork2RPC["ethereum"]
	}

	client, err := ethclient.DialContext(ctx, rpcUrl)
	if err != nil {
		return 0, err
	}
	address := common.HexToAddress(addressHex)
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return 0, err
	}

	contract := bind.NewBoundContract(address, parsed, client, client, client)

	out := uint8(0)
	err = contract.Call(&bind.CallOpts{}, &out, "decimals")
	return out, err
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
